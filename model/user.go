package model

type User struct {
	Id          int64  `gorm:"type:int;primary_key"`
	Name        string `gorm:"type:char(200)"`
	Email       string `gorm:"type:char(200)"`
	Description string `gorm:"type:varchar(500)"`
}
