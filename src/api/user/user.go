package user_router

import (
	"github.com/gin-gonic/gin"
)

type User_router struct{}

func NewUser() *User_router {
	return &User_router{}
}

type Info struct {
	Uname  string `form:"uname" json:"uname" binding:"required"`
	Passwd string `form:"passwd" json:"passwd" binding:"required"`
	Admin  bool   `form:"admin" json:"admin"`
}

func (s *User_router) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout",
	})
}
