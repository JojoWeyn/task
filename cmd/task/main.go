package main

import (
	"log"
	"task/internal/delivery/http"
	"task/internal/infrastructure/database"
	"task/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	database.AutoMigrate()
	// Репозитории
	userRepo := database.NewUserRepositoryDB(database.DB)
	refreshTokenRepo := database.NewRefreshTokenRepository(database.DB)

	// Бизнес-логика
	authUsecase := usecase.NewAuthUsecase(userRepo, refreshTokenRepo)

	// HTTP-обработчики
	authHandler := http.NewAuthHandler(*authUsecase)

	// Настройка маршрутов
	router := gin.Default()
	http.SetupRoutes(router, authHandler)

	log.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
