package repository

import (
	"time"

	"github.com/sreedharputtu/timeslothub-service/model"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Save(model.Booking) error
	FindByUserID(userID int64) ([]model.Booking, error)
	FindByCreatedBy(userID int64) ([]model.Booking, error)
	FindByUserIDAndBookingDate(userID int64, bookingDate time.Time) ([]model.Booking, error)
	FindByCalendarIDAndBookingDate(calendarID int64, bookingDate time.Time) ([]model.Booking, error)
}

type BookingsRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookingsRepository(DB *gorm.DB) BookingRepository {
	return &BookingsRepositoryImpl{DB: DB}
}

// FindByCalendarIDAndBookingDate implements BookingRepository.
func (bri *BookingsRepositoryImpl) FindByCalendarIDAndBookingDate(calendarID int64, bookingDate time.Time) ([]model.Booking, error) {
	var bookings []model.Booking
	result := bri.DB.Where("calendar_id=? and booking_date=?", calendarID, bookingDate).Find(&bookings)
	if result.Error != nil {
		return bookings, result.Error
	}
	return bookings, nil
}

// FindByCreatedBy implements BookingRepository.
func (bri *BookingsRepositoryImpl) FindByCreatedBy(userID int64) ([]model.Booking, error) {
	var bookings []model.Booking
	result := bri.DB.Where("created_by=? order by booking_date desc", userID).Find(&bookings)
	if result.Error != nil {
		return bookings, result.Error
	}
	return bookings, nil
}

// FindByUserID implements BookingRepository.
func (bri *BookingsRepositoryImpl) FindByUserID(userID int64) ([]model.Booking, error) {
	var bookings []model.Booking
	result := bri.DB.Where("user_id=? order by booking_date desc", userID).Find(&bookings)
	if result.Error != nil {
		return bookings, result.Error
	}
	return bookings, nil
}

// FindByUserIDAndBookingDate implements BookingRepository.
func (bri *BookingsRepositoryImpl) FindByUserIDAndBookingDate(userID int64, bookingDate time.Time) ([]model.Booking, error) {
	var bookings []model.Booking
	result := bri.DB.Where("user_id=? and booking_date=?", userID, bookingDate).Find(&bookings)
	if result.Error != nil {
		return bookings, result.Error
	}
	return bookings, nil
}

// Save implements BookingRepository.
func (bri *BookingsRepositoryImpl) Save(booking model.Booking) error {
	return bri.DB.Save(booking).Error
}
