package usecase

import (
	"errors"
	"strings"
	"time"

	"task/internal/entity"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your-secret-key")

func generateJWT(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (uint, error) {
	// Убедиться, что токен состоит из трех частей
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return 0, errors.New("invalid token format: token must have 3 parts")
	}

	// Разбор токена
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Убедиться, что используется правильный метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}

	// Проверка валидности токена
	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Извлечение user_id
	userID, ok := (*claims)["user_id"].(float64) // JWT стандартно конвертирует числа в float64
	if !ok {
		return 0, errors.New("invalid token claims: user_id not found")
	}

	// Проверка истечения срока действия (exp)
	if exp, ok := (*claims)["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return 0, errors.New("token has expired")
		}
	}

	return uint(userID), nil
}
