package controller

type CalendarController interface {
}

type calendarController struct{}

func NewCalendarController() CalendarController {
	return calendarController{}
}
