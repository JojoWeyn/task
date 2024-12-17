package main

import (
	"log"
	"task/internal/delivery/http"
	"task/internal/infrastructure/database"
	"task/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	database.AutoMigrate()

	userRepo := database.NewUserRepositoryDB(database.DB)
	projectRepo := database.NewProjectRepositoryDB(database.DB)
	refreshTokenRepo := database.NewRefreshTokenRepository(database.DB)

	authUsecase := usecase.NewAuthUsecase(userRepo, refreshTokenRepo)
	projectUsecase := usecase.NewProjectUsecase(projectRepo)

	projectHandler := http.NewProjectHandler(*projectUsecase)
	authHandler := http.NewAuthHandler(*authUsecase)

	router := gin.Default()
	http.AuthRoutes(router, authHandler)
	http.ProjectRoutes(router, projectHandler)
	http.ProfileRoutes(router, authHandler)

	log.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
