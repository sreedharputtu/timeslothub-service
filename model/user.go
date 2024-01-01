package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id          int64  `gorm:"type:int;primary_key"`
	Name        string `gorm:"type:char(200)"`
	Email       string `gorm:"type:char(200)"`
	Description string `gorm:"type:varchar(500)"`
	Password    string `gorm:"column:password_hash"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
