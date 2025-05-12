// internal/repository/user_repo.go
package repository

import (
	"github.com/user-service/internal/model"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitRepository(d *gorm.DB) {
	db = d
	db.AutoMigrate(&model.User{})
}

func CreateUser(u model.User) {
	db.Create(&u)
}

func GetUser(id string) *model.User {
	var u model.User
	result := db.First(&u, "id = ?", id)
	if result.Error != nil {
		return nil
	}
	return &u
}
