package challenge

import (
	challenge_ "src/logic/challenge"

	"github.com/gin-gonic/gin"
)

// 提交flag,进行检验
func (s *Challenge) FlagSubmit(c *gin.Context) {
	name := c.PostForm("name")
	flag := c.PostForm("flag")
	challenge_name := c.PostForm("challenge_name")
	var ch challenge_.Challenge
	ch.Name = challenge_name
	ch.Flags = flag
	err := ch.FlagSubmit(name)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
