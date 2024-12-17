package database

import (
	"errors"
	"task/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id uint) (*entity.User, error)
	isExist(email string) (bool, error)
}

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return &UserRepositoryDB{db: db}
}

func (r *UserRepositoryDB) Create(user *entity.User) error {
	if exist, err := r.isExist(user.Email); err != nil {
		return err
	} else if exist {
		return errors.New("пользователь с таким email уже существует")
	} else {
		return r.db.Create(user).Error
	}
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

func (r *UserRepositoryDB) isExist(email string) (bool, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
