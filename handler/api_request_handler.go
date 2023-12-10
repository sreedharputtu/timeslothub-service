package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sreedharputtu/timeslothub-service/model"
	"github.com/sreedharputtu/timeslothub-service/repository"
)

func NewRequestHandler(userRepository repository.UsersRepository) *RequestHandler {
	return &RequestHandler{userRespository: userRepository}
}

type RequestHandler struct {
	userRespository repository.UsersRepository
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
	c.HTML(201, "slot_settings_table.html", nil)
}

func (r *RequestHandler) UpdateCalenderSettings(c *gin.Context) {

}

func (r *RequestHandler) GetCalenderSettings(c *gin.Context) {

}
