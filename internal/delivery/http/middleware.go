package http

import (
	"net/http"
	"strings"
	"task/internal/usecase"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем заголовок Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Удаляем префикс "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Проверяем токен
		userID, err := usecase.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Сохраняем userID в контексте
		c.Set("user_id", userID)
		c.Next()
	}
}
