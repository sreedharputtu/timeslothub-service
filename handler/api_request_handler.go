package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/sreedharputtu/timeslothub-service/model"
	"github.com/sreedharputtu/timeslothub-service/repository"
)

const (
	timeformat = "^[0-2][0-3]:[0-5][0-9]+$"
)

func NewRequestHandler(userRepository repository.UsersRepository, ssr repository.SlotSettingsRepository) *RequestHandler {
	return &RequestHandler{userRespository: userRepository, slotSettingsRepository: ssr}
}

type RequestHandler struct {
	userRespository        repository.UsersRepository
	slotSettingsRepository repository.SlotSettingsRepository
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

func (r *RequestHandler) SaveCalendarSettings(c *gin.Context) {

}

func (r *RequestHandler) SaveSlotSettings(c *gin.Context) {
	dayOfWeek := c.Request.FormValue("day_of_week")
	startTime := c.Request.FormValue("start_time")
	endTime := c.Request.FormValue("end_time")

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
		DayOfWeek: dayOfWeek,
		StartTime: start,
		EndTime:   end,
		UserID:    int64(1),
	}
	err = r.slotSettingsRepository.Save(slotSettings)
	if err != nil {
		c.Status(500)
		return
	}

	c.HTML(201, "success_alert.html", gin.H{"msg": "slot added successfully"})

	//sslist, _ := r.slotSettingsRepository.FindAll()
	//ssdtolist := convertSlotSettings(sslist)

	// c.HTML(201, "slot_settings_table.html", gin.H{
	// 	"SlotSettingsList": ssdtolist,
	// })
}

func (r *RequestHandler) BookingsCalendar(c *gin.Context) {
	calendar := cal(0)
	fmt.Println(calendar)
	weekdayOrder := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	c.HTML(201, "bookings.html", gin.H{"Calendar": calendar, "Order": weekdayOrder})
}

func (r *RequestHandler) UpdateCalenderSettings(c *gin.Context) {

}

func (r *RequestHandler) GetCalenderSettings(c *gin.Context) {

}

func (r *RequestHandler) TimeSlots(c *gin.Context) {
	c.HTML(201, "timeslots.html", gin.H{})
}

type BookingsDays struct {
	WeekDay string
	Days    []int
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

	// Define a map to store days for each weekday
	daysPerWeekday := map[string][]int{
		"monday":    make([]int, 0),
		"tuesday":   make([]int, 0),
		"wednesday": make([]int, 0),
		"thursday":  make([]int, 0),
		"friday":    make([]int, 0),
		"saturday":  make([]int, 0),
		"sunday":    make([]int, 0),
	}

	// Calculate last day of current month
	lastDay := time.Date(currentYear, currentMonth+1, 0, 0, 0, 0, 0, time.UTC).AddDate(-1, 0, 0) // Subtract 1 day to get last day

	// Loop through each day of the month
	for day := 1; day <= lastDay.Day(); day++ { // Use lastDay.Day() instead of daysInMonth()
		// Get the weekday of the current day
		currentDay := time.Date(currentYear, currentMonth, day, 0, 0, 0, 0, time.UTC).Weekday()

		// Add the day to the corresponding weekday map
		switch currentDay {
		case time.Monday:
			daysPerWeekday["monday"] = append(daysPerWeekday["monday"], day)
			calendar[0].Days = append(calendar[0].Days, day)
		case time.Tuesday:
			daysPerWeekday["tuesday"] = append(daysPerWeekday["tuesday"], day)
			calendar[1].Days = append(calendar[1].Days, day)
		case time.Wednesday:
			daysPerWeekday["wednesday"] = append(daysPerWeekday["wednesday"], day)
			calendar[2].Days = append(calendar[2].Days, day)
		case time.Thursday:
			daysPerWeekday["thursday"] = append(daysPerWeekday["thursday"], day)
			calendar[3].Days = append(calendar[3].Days, day)
		case time.Friday:
			daysPerWeekday["friday"] = append(daysPerWeekday["friday"], day)
			calendar[4].Days = append(calendar[4].Days, day)
		case time.Saturday:
			daysPerWeekday["saturday"] = append(daysPerWeekday["saturday"], day)
			calendar[5].Days = append(calendar[5].Days, day)
		case time.Sunday:
			daysPerWeekday["sunday"] = append(daysPerWeekday["sunday"], day)
			calendar[6].Days = append(calendar[6].Days, day)
		}
	}

	fmt.Println("Days per weekday for current month:")
	for weekday, days := range daysPerWeekday {
		fmt.Printf("%-8s: %v\n", weekday, days)
	}

	return calendar
}
