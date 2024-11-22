package user

import "github.com/gin-gonic/gin"

func (s *User) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout",
	})
}
