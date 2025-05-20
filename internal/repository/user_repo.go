// internal/repository/user_repo.go
package repository

import (
	"github.com/user-service/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(u model.User)
	GetUser(id string) *model.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (c userRepository) CreateUser(u model.User) {
	c.db.Create(&u)
}

func (c userRepository) GetUser(id string) *model.User {
	var u model.User
	result := c.db.First(&u, "id = ?", id)
	if result.Error != nil {
		return nil
	}
	return &u
}
