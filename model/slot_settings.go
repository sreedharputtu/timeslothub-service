package model

type SlotSettings struct {
	ID        int64  `gorm:"type:int;primary_key"`
	DayOfWeek string `gorm:"type:char(20)"`
	StartTime int    `gorm:"type:int"`
	EndTime   int    `gorm:"type:int"`
	UserID    int64  `gorm:"type:int"`
}
