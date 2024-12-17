package usecase

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"task/internal/entity"
	"task/internal/infrastructure/database"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepo database.UserRepository
	jwtRepo  database.RefreshTokenRepository
}

func NewAuthUsecase(userRepo database.UserRepository, jwtRepo database.RefreshTokenRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo, jwtRepo: jwtRepo}
}

func generateRandomCode() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (u *AuthUsecase) Register(email, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		Email:          email,
		PasswordHash:   string(passwordHash),
		IsActive:       true,
		ActivationCode: generateRandomCode(),
		CreatedAt:      time.Now(),
	}
	if err := u.userRepo.Create(user); err != nil {
		return err
	}
	return nil
}

func (u *AuthUsecase) RegisterAsync(email, password string) <-chan error {
	result := make(chan error)

	go func() {
		defer close(result)
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			result <- err
			return
		}

		user := &entity.User{
			Email:          email,
			PasswordHash:   string(passwordHash),
			IsActive:       true,
			ActivationCode: generateRandomCode(),
			CreatedAt:      time.Now(),
		}
		if err := u.userRepo.Create(user); err != nil {
			result <- err
			return
		}
		result <- nil
	}()
	return result

}

func (u *AuthUsecase) Activate(email, code string) error {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return err
	}
	if user.ActivationCode != code {
		return errors.New("неверный код активации")
	}
	user.IsActive = true
	user.ActivationCode = ""
	return u.userRepo.Update(user)
}

func (u *AuthUsecase) GetUserByID(id uint) (*entity.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *AuthUsecase) Login(email, password string) (string, string, error) {
	user, err := u.userRepo.FindByEmail(email)

	if !user.IsActive {
		return "", "", errors.New("пользователь не активирован")
	}

	if err != nil {
		return "", "", errors.New("неверный пароль или логин")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", "", errors.New("неверный пароль или логин")
	}

	token, err := generateJWT(user)
	if err != nil {
		return "", "", err
	}

	refToken, err := u.generateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}
	return token, refToken, nil
}

func (u *AuthUsecase) generateRefreshToken(userID uint) (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}
	token := hex.EncodeToString(tokenBytes)

	refreshToken := &entity.RefreshToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		CreatedAt: time.Now(),
	}
	if err := u.jwtRepo.Create(refreshToken); err != nil {
		return "", err
	}

	return token, nil
}

func (u *AuthUsecase) RefreshAccessToken(refreshToken string) (string, string, error) {

	storedToken, err := u.jwtRepo.FindByToken(refreshToken)
	if err != nil || storedToken.ExpiresAt.Before(time.Now()) {
		return "", "", errors.New("недействительный или просроченный Refresh-токен")
	}

	user := &entity.User{ID: storedToken.UserID}
	accessToken, err := generateJWT(user)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := u.generateRefreshToken(storedToken.UserID)
	if err != nil {
		return "", "", err
	}

	if err := u.jwtRepo.Delete(refreshToken); err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}
