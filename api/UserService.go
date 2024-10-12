package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct{}

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Users []User

var user *User

func (us *UserService) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": User{},
	})
}

func (us *UserService) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": user,
	})
}
