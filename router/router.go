package router

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sreedharputtu/timeslothub-service/handler"
)

func NewRouter(rh *handler.RequestHandler) *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("./html/templates/**")
	r.Static("/images", "./html/images")
	r.StaticFS("/static", http.Dir("./static"))

	r.GET("/pages/calendars/list", rh.GetCalenderSettings)
	r.GET("/pages/bookings", rh.BookingsCalendar)

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

	protected := r.Group("")
	protected.Use(handler.AuthRequired())

	protected.GET("/", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		email := session.Get("user")
		fmt.Println("eamil", email)
		ctx.HTML(201, "index.html", gin.H{"Email": email})
	})

	///views/calendars/settings

	protected.GET("/pages/calendars/create", func(ctx *gin.Context) {
		ctx.HTML(201, "create_calendar.html", gin.H{})
	})

	protected.GET("/views/slots/settings", func(ctx *gin.Context) {
		ctx.HTML(201, "slot_settings.html", gin.H{})
	})

	protected.GET("/views/slots/bookings", rh.BookingsCalendar)

	protected.GET("/views/timeslots", rh.TimeSlots)

	protected.GET("/user_logout", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Clear()
		session.Save()
		ctx.HTML(200, "logout.html", nil)
	})

	r.GET("/login", func(ctx *gin.Context) {
		fmt.Println("inside login")
		ctx.HTML(201, "login.html", nil)
	})

	r.POST("/user/login", rh.Login)

	return r
}
