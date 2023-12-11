package model

import "time"

type SlotSettings struct {
	ID        int64     `gorm:"type:int;primary_key"`
	DayOfWeek string    `gorm:"type:char(20)"`
	StartTime time.Time `gorm:"type:time"`
	EndTime   time.Time `gorm:"type:time"`
	UserID    int64     `gorm:"type:int"`
}
