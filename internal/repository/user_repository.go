package repository

import "task/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id uint) (*entity.User, error)
}
