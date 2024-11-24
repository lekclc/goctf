package mid

import (
	"src/ctrl/utils"

	"github.com/gin-gonic/gin"
)

func AuthMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		Jwt := utils.GetJwt(c)
		ca, err := Jwt.GetClaims()

		if err != nil || ca.Username != name {
			c.JSON(401, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		Jwt := utils.GetJwt(c)
		ca, err := Jwt.GetClaims()
		if err != nil {
			c.JSON(401, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		if !ca.Auth {
			c.JSON(401, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
