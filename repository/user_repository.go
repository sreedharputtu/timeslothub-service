package repository

import (
	"github.com/sreedharputtu/timeslothub-service/model"
	"gorm.io/gorm"
)

type UsersRepository interface {
	Save(user model.User) error
	FindByEmail(email string) (*model.User, error)
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

func (ur *UsersRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := ur.Db.Where("email=?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UsersRepositoryImpl) FindAll() ([]model.User, error) {
	var users []model.User
	result := ur.Db.Find(&users)
	return users, result.Error
}
