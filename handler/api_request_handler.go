package handler

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

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

	sslist, _ := r.slotSettingsRepository.FindAll()
	ssdtolist := convertSlotSettings(sslist)

	c.HTML(201, "slot_settings_table.html", gin.H{
		"SlotSettingsList": ssdtolist,
	})
}

func (r *RequestHandler) UpdateCalenderSettings(c *gin.Context) {

}

func (r *RequestHandler) GetCalenderSettings(c *gin.Context) {

}
