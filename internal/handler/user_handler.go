package handler

import (
	"net/http"

	"github.com/user-service/internal/dto"
	"github.com/user-service/internal/model"
	"github.com/user-service/internal/service"
	logger "github.com/user-service/log"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (s *UserHandler) CreateUser(c *gin.Context) {
	var u dto.ReqUserDto
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Log.Infof("Request: %v", u)
	err := s.service.CreateUser(model.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusCreated, nil)
}

func (s *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user := s.service.GetUser(id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	resp := dto.RespUserDto{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	c.JSON(http.StatusOK, resp)
}

// func fetchUser(c *gin.Context){
// 	user := service.FetchUser()
// }
