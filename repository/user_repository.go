package repository

import (
	"github.com/sreedharputtu/timeslothub-service/model"
	"gorm.io/gorm"
)

type UsersRepository interface {
	Save(user model.User) error
}

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

func (ur *UsersRepositoryImpl) Save(user model.User) error {
	result := ur.Db.Create(&user)
	return result.Error
}

func (ur *UsersRepositoryImpl) FindAll() ([]model.User, error) {
	var users []model.User
	result := ur.Db.Find(&users)
	return users, result.Error
}
