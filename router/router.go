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

	r.GET("/pages/my/calendars/list", rh.GetMyCalenders)

	protected := r.Group("")
	protected.Use(handler.AuthRequired())
	protected.GET("/", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		email := session.Get("user_email")
		ctx.HTML(201, "index.html", gin.H{"Email": email})
	})

	//apis
	protected.POST("/api/v1/my/calendars", rh.SaveMyCalendar)
	protected.GET("/api/v1/my/calendars", rh.GetMyCalenders)
	protected.GET("/api/v1/calendars", rh.GetCalenders)
	protected.GET("/api/v1/calendars/:calendar_id/slots", rh.GetSlotsByCalendarID)
	protected.POST("/api/v1/slots", rh.SaveSlot)

	//pages
	protected.GET("/page/bookings", rh.BookingsCalendar)
	protected.GET("/page/my/calendars/create", func(ctx *gin.Context) {
		ctx.HTML(201, "create_calendar.html", gin.H{})
	})
	protected.GET("/page/slots/settings", func(ctx *gin.Context) {
		ctx.HTML(201, "slot_settings.html", gin.H{})
	})
	protected.GET("/page/slots/bookings", rh.BookingsCalendar)

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

	r.GET("/register", func(ctx *gin.Context) {
		fmt.Println("inside register")
		ctx.HTML(201, "register.html", nil)
	})

	r.POST("/user/login", rh.Login)
	r.POST("/user/register", rh.Register)

	r.GET("/bookings", rh.GetBookings)
	r.POST("/bookings", rh.SaveBooking)
	r.GET("/my/bookings", rh.GetMyBookings)
	r.GET("/bookings/received", rh.GetReceivedBookings)

	return r
}
