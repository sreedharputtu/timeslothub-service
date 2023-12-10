package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sreedharputtu/timeslothub-service/repository"
)

type UIRequestHandler struct {
	usersRespository repository.UsersRepository
}

func NewUIRequestHandler(usersRepository repository.UsersRepository) *UIRequestHandler {
	return &UIRequestHandler{usersRespository: usersRepository}
}

func (ui *UIRequestHandler) Login(c *gin.Context) {

}

func (ui *UIRequestHandler) Register(c *gin.Context) {

}

func (ui *UIRequestHandler) CalendarSettings(c *gin.Context) {

}

func (ui *UIRequestHandler) SlotSettings(c *gin.Context) {

}

func (ui *UIRequestHandler) Calendar(c *gin.Context) {

}
