package service

import (
	"github.com/user-service/internal/model"
	"github.com/user-service/internal/repository"
	"github.com/user-service/pkg"
)

type UserService interface {
	CreateUser(u model.User) error
	GetUser(id string) *model.User
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (c *userService) CreateUser(u model.User) error {
	hash, err := pkg.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	c.repo.CreateUser(u)
	return nil
}

func (c *userService) GetUser(id string) *model.User {
	return c.repo.GetUser(id)
}
