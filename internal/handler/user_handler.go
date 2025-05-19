package handler

import (
	"net/http"

	"github.com/user-service/internal/dto"
	"github.com/user-service/internal/model"
	"github.com/user-service/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	u := r.Group("/users")
	{
		u.POST("/", createUser)
		u.GET("/:id", getUser)
	}
}

func createUser(c *gin.Context) {
	var u dto.ReqUserDto
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.CreateUser(model.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusCreated, nil)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	user := service.GetUser(id)
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
