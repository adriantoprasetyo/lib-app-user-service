package service

import (
	"github.com/user-service/internal/model"
	"github.com/user-service/internal/repository"
	"github.com/user-service/pkg"
)

func CreateUser(u model.User) error {
	hash, err := pkg.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	repository.CreateUser(u)
	return nil
}

func GetUser(id string) *model.User {
	return repository.GetUser(id)
}
