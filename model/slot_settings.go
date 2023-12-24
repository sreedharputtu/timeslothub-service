package model

import "time"

type SlotSettings struct {
	ID          int64     `gorm:"type:int;primary_key"`
	DayOfWeekID int       `gorm:"type:int"`
	DayOfWeek   string    `gorm:"type:char(20)"`
	StartTime   int       `gorm:"type:int"`
	EndTime     int       `gorm:"type:int"`
	UserID      int64     `gorm:"type:int"`
	CalendarID  int64     `gorm:"type:int"`
	CreatedAt   time.Time `gorm:"type:timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp"`
}
