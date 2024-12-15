package database

import (
	"task/internal/entity"
	"task/internal/repository"

	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryDB{db: db}
}

func (r *UserRepositoryDB) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryDB) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepositoryDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryDB) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
