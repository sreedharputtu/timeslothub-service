package controller

type BookingsController interface {
}

type bookingsController struct {
}

func NewBookingsController() BookingsController {
	return &bookingsController{}
}
