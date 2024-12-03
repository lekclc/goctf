package user

import (
	"net/http"
	user_ "src/logic/user"

	"github.com/gin-gonic/gin"
)

// 注册
func (s *User) Register(c *gin.Context) {
	name := c.PostForm("name")
	passwd := c.PostForm("passwd")
	u := user_.GetUser(name, passwd)
	is, err := u.Register()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	if is {
		c.JSON(200, gin.H{
			"message": "register",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "error",
	})
}
