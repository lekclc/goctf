package mid

import (
	"src/ctrl/utils"

	"github.com/gin-gonic/gin"
)

func AuthMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		token := c.PostForm("token")
		Jwt := utils.GetJwt(token)
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
		name := c.PostForm("name")
		token := c.PostForm("token")
		Jwt := utils.GetJwt(token)
		ca, err := Jwt.GetClaims()
		if err != nil || ca.Username != name {
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
