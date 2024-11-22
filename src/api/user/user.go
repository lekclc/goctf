package user

import (
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

type Info struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" binding:"required"`
	Admin    bool   `form:"admin" json:"admin"`
}

func (s *User) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout",
	})
}
