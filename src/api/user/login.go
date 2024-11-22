package user

import (
	"net/http"
	"src/ctrl/utils"
	user_ "src/logic/user"

	"github.com/gin-gonic/gin"
)

func (s *User) Login(c *gin.Context) {
	var info Info
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	}
	u := user_.NewUser(info.Username, info.Passwd)
	is, isadmin, err := u.Login()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	if !is {
		c.JSON(http.StatusOK, gin.H{
			"message": "login failed",
		})
		return
	}
	token, err := utils.GetJwt(c).Get(isadmin, info.Username)
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
