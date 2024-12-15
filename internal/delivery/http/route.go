package http

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, authHandler *AuthHandler) {
	// Public routes
	router.POST("auth/register", authHandler.Register)
	router.POST("auth/login", authHandler.Login)

	// Protected routes
	protected := router.Group("/")
	protected.Use(AuthMiddleware())

	// Пример защищённого маршрута
	protected.GET("/profile", authHandler.GetProfile)

}
