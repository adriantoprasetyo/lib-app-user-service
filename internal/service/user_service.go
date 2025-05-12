package service

import (
	"github.com/user-service/internal/model"
	"github.com/user-service/internal/repository"
)

func CreateUser(u model.User) {
	repository.CreateUser(u)
}

func GetUser(id string) *model.User {
	return repository.GetUser(id)
}
