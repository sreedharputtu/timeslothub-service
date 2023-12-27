package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sreedharputtu/timeslothub-service/handler"
)

func NewRouter(rh *handler.RequestHandler) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("./html/templates/**")

	r.Static("/images", "./html/images")
	r.StaticFS("/static", http.Dir("./static"))

	rg := r.Group("/api/v1")
	rg.POST("/users", rh.SaveUser)
	rg.PUT("/users/:user_id")
	rg.POST("/calendars/settings", rh.SaveCalendarSettings)
	rg.PUT("/calendars/settings/:user_id", rh.UpdateCalenderSettings)
	rg.GET("/calendars/settings", rh.GetCalenderSettings)
	rg.GET("/calendars/:calendar_id/slots", rh.GetSlotSettingsByCalendarID)
	rg.POST("/slots/settings", rh.SaveSlotSettings)
	rg.GET("/slots/settings/:user_id", rh.GetCalenderSettings)
	rg.PUT("/slots/settings/:slot_id", rh.GetCalenderSettings)

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(201, "index.html", gin.H{})
	})

	///views/calendars/settings

	r.GET("/views/calendars/settings", func(ctx *gin.Context) {
		ctx.HTML(201, "create_calendar.html", gin.H{})
	})

	r.GET("/views/slots/settings", func(ctx *gin.Context) {
		ctx.HTML(201, "slot_settings.html", gin.H{})
	})

	r.GET("/views/slots/bookings", rh.BookingsCalendar)

	r.GET("/views/timeslots", rh.TimeSlots)

	return r
}
