package user

import (
	user_ "src/logic/user"

	"github.com/gin-gonic/gin"
)

func (s *User) Del(c *gin.Context) {
	name := c.PostForm("name")
	var u user_.User
	err := u.Del(name)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})

}
