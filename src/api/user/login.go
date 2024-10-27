package user_router

import (
	"net/http"
	user_ "src/logic/user"

	"github.com/gin-gonic/gin"
)

func (s *User_router) Login(c *gin.Context) {
	var info Info
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	}
	u := user_.NewUser(info.Uname, info.Passwd)
	is, err := u.Login()
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
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}
