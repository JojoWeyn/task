package http

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.Engine, authHandler *AuthHandler) {

	authGroup := router.Group("/auth")

	authGroup.POST("/register", authHandler.RegisterAsync)
	authGroup.POST("/login", authHandler.Login)
}

func ProfileRoutes(router *gin.Engine, authHandler *AuthHandler) {
	protected := router.Group("/profile")
	protected.Use(AuthMiddleware())

	protected.GET("/", authHandler.GetProfile)

}

func ProjectRoutes(router *gin.Engine, projectHandler *ProjectHandler) {
	projectGroup := router.Group("/project")
	projectGroup.Use(AuthMiddleware())
	projectGroup.POST("/create", projectHandler.CreateProject)
	projectGroup.GET("/list/", projectHandler.ListByUser)
	projectGroup.GET("/:id", projectHandler.FindByID)
}
