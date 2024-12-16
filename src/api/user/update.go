package user

import (
	"net/http"
	user_ "src/logic/user"

	"github.com/gin-gonic/gin"
)

func (s *User) Update(c *gin.Context) {
	name := c.PostForm("name")
	passwd := c.PostForm("passwd")
	u := user_.GetUser(name, passwd)
	err := u.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "update success",
	})
}
