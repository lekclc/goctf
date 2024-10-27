package user_router

import (
	"net/http"
	user_ "src/logic/user"

	"github.com/gin-gonic/gin"
)

func (s *User_router) Register(c *gin.Context) {
	var info Info
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	u := user_.NewUser(info.Uname, info.Passwd)
	is, err := u.Register()
	if err != nil {
		c.JSON(200, gin.H{
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
	c.JSON(200, gin.H{
		"message": "error",
	})
}
