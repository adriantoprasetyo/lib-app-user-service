package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/user-service/internal/config"
	"github.com/user-service/internal/handler"
	"github.com/user-service/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// load
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	cfg := config.Load()
	fmt.Println("data", os.Getenv("DB_HOST"))

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	repository.InitRepository(db)

	r := gin.Default()
	handler.RegisterRoutes(r)
	log.Fatal(r.Run(":" + cfg.DBPort))
}
