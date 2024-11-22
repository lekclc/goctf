package user

import (
	"net/http"
	user_ "src/logic/user"

	"github.com/gin-gonic/gin"
)

func (s *User) Register(c *gin.Context) {
	var info Info
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	u := user_.GetUser(info.Username, info.Passwd)
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
