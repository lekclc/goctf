package user

import "github.com/gin-gonic/gin"

type User struct{}

func NewUser() *User {
	return &User{}
}

func (s *User) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func (s *User) Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register",
	})
}

func (s *User) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout",
	})
}