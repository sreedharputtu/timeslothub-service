package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/sreedharputtu/timeslothub-service/model"
	"github.com/sreedharputtu/timeslothub-service/repository"
)

const (
	timeformat = "^[0-2][0-3]:[0-5][0-9]+$"
)

func NewRequestHandler(userRepository repository.UsersRepository, ssr repository.SlotSettingsRepository, cri repository.CalendarSettingsRepository, br repository.BookingRepository) *RequestHandler {
	return &RequestHandler{userRespository: userRepository, slotSettingsRepository: ssr, calendarRepo: cri, br: br}
}

type RequestHandler struct {
	userRespository        repository.UsersRepository
	slotSettingsRepository repository.SlotSettingsRepository
	calendarRepo           repository.CalendarSettingsRepository
	br                     repository.BookingRepository
}

func Health(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (r *RequestHandler) SaveUser(c *gin.Context) {
	user := model.User{
		Name:        c.Request.FormValue("name"),
		Email:       c.Request.FormValue("email"),
		Description: c.Request.FormValue("description"),
	}
	err := r.userRespository.Save(user)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.HTML(201, "bookings.html", nil)
}

func (r *RequestHandler) SaveMyCalendar(c *gin.Context) {

	slotTime, err := strconv.Atoi(c.Request.FormValue("slot_time"))
	if err != nil {
		log.Error(err)
		c.Status(400)
		return
	}

	autoAccept := false
	autoAcceptStr := c.Request.FormValue("auto_accept")
	if autoAcceptStr == "on" {
		autoAccept = true
	}

	utcOffsetStr := c.Request.FormValue("local_time_offset")
	utcOffset, err := strconv.Atoi(utcOffsetStr)

	if err != nil {
		c.Status(400)
		return
	}

	session := sessions.Default(c)
	userIDRaw := session.Get("user_id")
	if userIDRaw == nil {
		c.Status(500)
		return
	}

	userID := userIDRaw.(int64)

	calendar := model.CalendarSettings{
		CalendarName: c.Request.FormValue("calendar_name"),
		SlotTime:     int32(slotTime),
		AutoAccept:   autoAccept,
		UTCOffset:    utcOffset,
		UserID:       userID,
		CreatedAt:    time.Now(),
	}

	calendarID, err := r.calendarRepo.Save(calendar)
	c.HTML(201, "add_slot.html", gin.H{
		"calendar_name": calendar.CalendarName,
		"calendar_id":   calendarID,
	})
}

func (r *RequestHandler) SaveSlot(c *gin.Context) {
	dayOfWeek := c.Request.FormValue("day_of_week")
	startTime := c.Request.FormValue("start_time")
	endTime := c.Request.FormValue("end_time")
	calendarIDStr := c.Request.FormValue("calendar_id")
	log.Debug(fmt.Sprintf("slot is creating for  calendar id:%s", calendarIDStr))

	calendarID, _ := strconv.ParseInt(calendarIDStr, 10, 64)

	match, err := regexp.MatchString(timeformat, startTime)
	if err != nil || !match {
		log.Error(err)
		c.Status(400)
		return
	}

	match, err = regexp.MatchString(timeformat, endTime)
	if err != nil || !match {
		log.Error(err)
		c.Status(400)
		return
	}

	start, _ := strconv.Atoi(strings.Replace(startTime, ":", "", -1))
	end, _ := strconv.Atoi(strings.Replace(endTime, ":", "", -1))

	if start > end {
		log.Error("compare failure")
		c.Status(400)
		return
	}

	log.Debug("day_of_week:", dayOfWeek)

	slotSettings := model.SlotSettings{
		DayOfWeek:  dayOfWeek,
		StartTime:  start,
		EndTime:    end,
		UserID:     int64(1),
		CalendarID: calendarID,
	}
	err = r.slotSettingsRepository.Save(slotSettings)
	if err != nil {
		c.Status(500)
		return
	}

	//c.HTML(201, "success_alert.html", gin.H{"msg": "slot added successfully"})

	sslist, _ := r.slotSettingsRepository.FindByCalendarID(calendarID)
	ssdtolist := convertSlotSettings(sslist)

	c.HTML(201, "slot_settings_table.html", gin.H{
		"SlotSettingsList": ssdtolist,
		"SlotAddStatus":    true,
		"Msg":              "slot added successfully",
	})
}

func (r *RequestHandler) BookingsCalendar(c *gin.Context) {
	calendar := cal(0)
	weekdayOrder := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	c.HTML(201, "bookings_select_date.html", gin.H{"Calendar": calendar, "Order": weekdayOrder})
}

func (r *RequestHandler) UpdateCalenderSettings(c *gin.Context) {

}

func (r *RequestHandler) GetMyCalenders(c *gin.Context) {
	session := sessions.Default(c)
	userIDRaw := session.Get("user_id")
	userID := int64(0)
	if userIDRaw != nil {
		userID = userIDRaw.(int64)
	}
	calendars, err := r.calendarRepo.FindByUserID(userID)
	if err != nil {
		log.Error(err)
		return
	}
	c.HTML(200, "view_calendars_list.html", gin.H{
		"Calendars": calendars,
	})
}

func (r *RequestHandler) GetCalenders(c *gin.Context) {
	emailParam := c.Query("user_email")
	calendarIDParam := c.Query("calendar_id")

	log.Debug(emailParam)
	log.Debug(calendarIDParam)

	if emailParam == "" && calendarIDParam == "" {
		return
	}

	currentCalendar := cal(0)

	if emailParam != "" {
		//validate email
		user, err := r.userRespository.FindByEmail(emailParam)
		if err != nil {
			log.Error(err)
			return
		}
		calendars, err := r.calendarRepo.FindByUserID(user.Id)
		if err != nil {
			log.Error(err)
			return
		}

		c.HTML(200, "calendars.html", gin.H{
			"Calendars":       calendars,
			"CurrentCalendar": currentCalendar,
			"CurrentMonth":    "January",
		})

		return
	}

	var calendarID int64
	if calendarIDParam != "" {
		calendarIDInternal, err := strconv.ParseInt(calendarIDParam, 10, 64)
		if err != nil {
			return
		}
		calendarID = calendarIDInternal
	}

	calendar, err := r.calendarRepo.FindByID(calendarID)
	if err != nil {
		return
	}

	calendars := make([]model.CalendarSettings, 1)
	calendars[0] = calendar
	c.HTML(200, "view_calendars_list.html", gin.H{
		"Calendars":       calendars,
		"CurrentCalendar": currentCalendar,
	})
	return
}

func (r *RequestHandler) GetSlotsByCalendarID(c *gin.Context) {
	calendarID, err := strconv.ParseInt(c.Param("calendar_id"), 10, 64)
	if err != nil {
		log.Error(err)
		return
	}
	slots, err := r.slotSettingsRepository.FindByCalendarID(calendarID)
	if err != nil {
		log.Error(err)
		return
	}

	c.HTML(201, "slot_settings_table.html", gin.H{
		"SlotSettingsList": slots,
	})

}

type Timeslot struct {
	StartTime int16
	EndTime   int16
	IsBooked  bool
}

func (r *RequestHandler) TimeSlots(c *gin.Context) {
	timeslots := make([]Timeslot, 3)

	timeslot1 := Timeslot{
		0000,
		1200,
		false,
	}

	timeslot2 := Timeslot{
		1200,
		1400,
		true,
	}

	timeslot3 := Timeslot{
		1400,
		1600,
		false,
	}

	timeslots[0] = timeslot1
	timeslots[1] = timeslot2
	timeslots[2] = timeslot3

	c.HTML(201, "bookings_select_timeslot.html", gin.H{
		"Timeslots": timeslots,
	})
}

type BookingsDays struct {
	WeekDay string
	Days    []Day
	Month   int
	Year    int
}

type Day struct {
	Day  int
	Date string
}

func cal(month int) []BookingsDays {
	// Get current month and year
	now := time.Now()
	currentMonth := now.Month()
	currentYear := now.Year()

	ddd := map[int]string{
		0: "monday",
		1: "tuesday",
		2: "wednesday",
		3: "thursday",
		4: "friday",
		5: "saturday",
		6: "sunday",
	}

	calendar := make([]BookingsDays, 7)
	for i := range calendar {
		calendar[i] = BookingsDays{
			WeekDay: ddd[i],
		}
	}

	// Calculate last day of current month
	lastDay := time.Date(currentYear, currentMonth+1, 0, 0, 0, 0, 0, time.UTC).AddDate(-1, 0, 0) // Subtract 1 day to get last day

	for day := 1; day <= lastDay.Day(); day++ { // Use lastDay.Day() instead of daysInMonth()
		currentDay := time.Date(currentYear, currentMonth, day, 0, 0, 0, 0, time.UTC).Weekday()
		switch currentDay {
		case time.Monday:
			calendar[0].Days = append(calendar[0].Days, Day{
				Day:  day,
				Date: fmt.Sprintf("%d/%d/%d", day, currentMonth, currentYear),
			})
		case time.Tuesday:
			calendar[1].Days = append(calendar[1].Days, Day{
				Day:  day,
				Date: fmt.Sprintf("%d/%d/%d", day, currentMonth, currentYear),
			})
		case time.Wednesday:
			calendar[2].Days = append(calendar[2].Days, Day{
				Day:  day,
				Date: fmt.Sprintf("%d/%d/%d", day, currentMonth, currentYear),
			})
		case time.Thursday:
			calendar[3].Days = append(calendar[3].Days, Day{
				Day:  day,
				Date: fmt.Sprintf("%d/%d/%d", day, currentMonth, currentYear),
			})
		case time.Friday:
			calendar[4].Days = append(calendar[4].Days, Day{
				Day:  day,
				Date: fmt.Sprintf("%d/%d/%d", day, currentMonth, currentYear),
			})
		case time.Saturday:
			calendar[5].Days = append(calendar[5].Days, Day{
				Day:  day,
				Date: fmt.Sprintf("%d/%d/%d", day, currentMonth, currentYear),
			})
		case time.Sunday:
			calendar[6].Days = append(calendar[6].Days, Day{
				Day:  day,
				Date: fmt.Sprintf("%d/%d/%d", day, currentMonth, currentYear),
			})
		}
	}
	return calendar
}

func (rh *RequestHandler) Login(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	if email == "" || password == "" {
		log.Error("email or password empty")
		c.JSON(401, gin.H{
			"error": "invalid credentails",
		})
		return
	}

	user, err := rh.userRespository.FindByEmail(email)
	if err != nil {
		log.Error(err)
		c.JSON(401, gin.H{
			"error": "could not find user with given email",
		})
		return
	}

	if email == "sreedharputtu@gmail.com" {
		user.HashPassword("123456789")
	}

	if email == "swathiputtu@gmail.com" {
		user.HashPassword("912345678")
	}

	err = user.CheckPassword(password)
	if err != nil {
		log.Error(err)
		c.JSON(401, gin.H{
			"error": "invalid password",
		})
		return
	}

	jwtWrapper := JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	session := sessions.Default(c)
	session.Set("state", signedToken)
	session.Set("user_id", user.Id)
	session.Set("user_email", user.Email)
	err = session.Save()
	if err != nil {
		log.Error("error while saving session", err)
	}
	// refreshToken, err := jwtWrapper.RefreshToken(user.Email)
	// if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"Error": "Error Signing Token",
	// 	})
	// 	c.Abort()
	// 	return
	// }
	c.Redirect(http.StatusFound, "/")

}

type BookingSlotDTO struct {
	ID         int64
	CalendarID int64
	StartTime  string
	EndTime    string
	Status     string
}

func convertBookingsModelToDTO(bookings []model.Booking) []BookingSlotDTO {
	var dtos []BookingSlotDTO
	for _, booking := range bookings {
		dtos = append(dtos, BookingSlotDTO{
			ID: booking.ID,
		})
	}
	return dtos
}

func (rh *RequestHandler) GetBookings(c *gin.Context) {
	selectedDateParam := c.Query("selected_date")
	selectedDayParam := c.Query("selected_day")
	calendarIDParam := c.Query("calendar_id")

	calendarID, err := strconv.ParseInt(calendarIDParam, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	selectedCalendar, err := rh.calendarRepo.FindByID(calendarID)
	if err != nil {
		c.Status(500)
		return
	}
	// Fetch slots
	slots, err := rh.slotSettingsRepository.FindByCalendarID(calendarID)
	if err != nil {
		c.Status(500)
		return
	}
	var selectedSlots []model.SlotSettings
	for _, slot := range slots {
		if strings.EqualFold(slot.DayOfWeek, selectedDayParam) {
			selectedSlots = append(selectedSlots, slot)
		}
	}

	//selectedCalendar.SlotTime
	//create bookings

	log.Error(selectedDateParam)
	log.Error(calendarIDParam)
	bookings := make([]model.Booking, 2)
	bookings[0] = model.Booking{ID: 1, StartDateTime: time.Now(), EndDateTime: time.Now()}
	bookings[1] = model.Booking{ID: 2}
	c.JSON(200, bookings)
}

func (rh *RequestHandler) SaveBooking(c *gin.Context) {
	var req CreateBookingRequestDTO
	err := c.BindJSON(&req)
	if err != nil {
		log.Error(err)
		c.Status(400)
		return
	}

	err = rh.br.Save(toBookingModel(req, int64(1)))
	if err != nil {
		log.Error(err)
		c.Status(500)
		return
	}

}
