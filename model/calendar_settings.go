package model

import "time"

type CalendarSettings struct {
	ID           int64     `gorm:"type:int;primary_key"`
	CalendarName string    `gorm:"type:char(200)"`
	SlotTime     int32     `gorm:"type:int"`
	AutoAccept   bool      `gorm:"type:bool"`
	UserID       int64     `gorm:"type:int"`
	CreatedAt    time.Time `gorm:"type:timestamp"`
	UpdatedAt    time.Time `gorm:"type:timestamp"`
}
