package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sreedharputtu/timeslothub-service/handler"
)

func NewRouter(rh *handler.RequestHandler) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("./html/templates/**")

	r.Static("/images", "./html/images")

	rg := r.Group("/api/v1")
	rg.POST("/users", rh.SaveUser)
	rg.PUT("/users/:user_id")
	rg.POST("/calendars/settings/:user_id", rh.SaveCalendarSettings)
	rg.PUT("/calendars/settings/:user_id", rh.UpdateCalenderSettings)
	rg.GET("/calendars/settings/:user_id", rh.GetCalenderSettings)
	rg.POST("/slots/settings", rh.SaveSlotSettings)
	rg.GET("/slots/settings/:user_id", rh.GetCalenderSettings)
	rg.PUT("/slots/settings/:slot_id", rh.GetCalenderSettings)

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(201, "index.html", gin.H{})
	})

	///views/calendars/settings

	r.GET("/views/calendars/settings", func(ctx *gin.Context) {
		ctx.HTML(201, "calendar_settings.html", gin.H{})
	})

	r.GET("/views/slots/settings", func(ctx *gin.Context) {
		ctx.HTML(201, "slot_settings.html", gin.H{})
	})

	r.GET("/views/slots/bookings", func(ctx *gin.Context) {
		ctx.HTML(201, "bookings.html", gin.H{})
	})

	return r
}
