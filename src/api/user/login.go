package user

import (
	"net/http"
	"src/ctrl/utils"
	user_ "src/logic/user"

	"github.com/gin-gonic/gin"
)

// 登录
func (s *User) Login(c *gin.Context) {
	name := c.PostForm("name")
	passwd := c.PostForm("passwd")
	u := user_.GetUser(name, passwd)
	is, isadmin, err := u.Login()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	if !is {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "login failed",
		})
		return
	}
	token, err := utils.GetJwt("").Get(isadmin, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
		"token":   token,
	})
}
