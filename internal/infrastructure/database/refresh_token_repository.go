package database

import (
	"task/internal/entity"

	"gorm.io/gorm"
)

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return NewRefreshTokenRepositoryDB(db)
}

type RefreshTokenRepository interface {
	Create(token *entity.RefreshToken) error
	FindByToken(token string) (*entity.RefreshToken, error)
	Delete(token string) error
}

type RefreshTokenRepositoryDB struct {
	db *gorm.DB
}

func NewRefreshTokenRepositoryDB(db *gorm.DB) RefreshTokenRepository {
	return &RefreshTokenRepositoryDB{db: db}
}

func (r *RefreshTokenRepositoryDB) Create(token *entity.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *RefreshTokenRepositoryDB) FindByToken(token string) (*entity.RefreshToken, error) {
	var refreshToken entity.RefreshToken
	if err := r.db.Where("token = ?", token).First(&refreshToken).Error; err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *RefreshTokenRepositoryDB) Delete(token string) error {
	return r.db.Where("token = ?", token).Delete(&entity.RefreshToken{}).Error
}
