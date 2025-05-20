package infrastructure

import (
	"github.com/user-service/internal/repository"
	"github.com/user-service/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	UserService service.UserService
}

func InitApp(dsn string) (*App, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(db)

	return &App{
		UserService: service.NewUserService(userRepo),
	}, nil
}
