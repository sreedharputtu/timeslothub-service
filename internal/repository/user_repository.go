package repository

import "github.com/sreedharputtu/timeslothub-service/internal/model"

type UsersRepository interface {
	Save(user model.User) error
}

type UsersRepositoryImpl struct {
}

func NewUserRepository() UsersRepository {
	return &UsersRepositoryImpl{}
}

func (ur *UsersRepositoryImpl) Save(user model.User) error {
	return nil
}
