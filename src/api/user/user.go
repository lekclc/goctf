package user_router

import (
	"net/http"
	user_ "src/logic/user"

	"github.com/gin-gonic/gin"
)

type User_router struct{}

func NewUser() *User_router {
	return &User_router{}
}

type Info struct {
	uname  string
	passwd string
}

func (s *User_router) Login(c *gin.Context) {
	var info Info
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	}
	u := user_.NewUser(info.uname, info.passwd)
	is, err := u.Login()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	}
	if !is {
		c.JSON(http.StatusOK, gin.H{
			"message": "login filed",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}

func (s *User_router) Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register",
	})
}

func (s *User_router) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout",
	})
}
