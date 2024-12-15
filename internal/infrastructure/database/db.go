package database

import (
	"log"

	"task/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=admin dbname=piska1 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection successful!")
}

func AutoMigrate() {
	DB.AutoMigrate(&entity.User{})
	DB.AutoMigrate(&entity.RefreshToken{})
	DB.AutoMigrate(&entity.Project{})
	DB.AutoMigrate(&entity.Task{})
	DB.AutoMigrate(&entity.ProjectGroup{})

}
