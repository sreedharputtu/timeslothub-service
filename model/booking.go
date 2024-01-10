package model

import "time"

type Booking struct {
	ID            int64     `gorm:"type:int;primary_key"`
	UserID        int64     `gorm:"type:int"`
	CalendarID    int64     `gorm:"type:int"`
	Status        string    `gorm:"type:char(100)"`
	BookingDate   time.Time `gorm:"type:timestamp"`
	StartDateTime time.Time `gorm:"type:timestamp"`
	EndDateTime   time.Time `gorm:"type:timestamp"`
	CreatedBy     int64     `gorm:"type:int"`
	CreatedAt     time.Time `gorm:"type:timestamp"`
	UpdatedAt     time.Time `gorm:"type:timestamp"`
}
