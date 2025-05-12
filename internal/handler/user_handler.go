package handler

import (
	"net/http"

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
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.CreateUser(u)
	c.JSON(http.StatusCreated, u)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	user := service.GetUser(id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
