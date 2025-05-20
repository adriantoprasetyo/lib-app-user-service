package main

import (
	"fmt"
	"log"
	"os"

	"github.com/user-service/internal/config"
	"github.com/user-service/internal/handler"
	infra "github.com/user-service/internal/infrastructure"
	"github.com/user-service/internal/middleware"
	logger "github.com/user-service/log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	fmt.Println("data", os.Getenv("DB_HOST"))

	app, err := infra.InitApp(cfg.DSN())
	fmt.Println("APP:", app.UserService)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	logger.Init()
	r := gin.New()
	r.Use(middleware.ZapLogger())

	userHandler := handler.NewUserHandler(app.UserService)
	u := r.Group("/users")
	{
		u.POST("/", userHandler.CreateUser)
		u.GET("/:id", userHandler.GetUser)
	}
	log.Fatal(r.Run(":" + cfg.ServicePort))
}
