package challenge

import (
	challenge_ "src/logic/challenge"

	"github.com/gin-gonic/gin"
)

// 添加flag
func (s *Challenge) AddFlag(c *gin.Context) {
	challenge_name := c.PostForm("challenge_name")
	flags := c.PostForm("flags")
	var ch challenge_.Challenge
	ch.Name = challenge_name
	ch.Flags = flags
	status := ch.AddFlag()
	if status == 1 {
		c.JSON(400, gin.H{
			"message": "题目不存在",
		})
		return
	} else if status == 2 {
		c.JSON(400, gin.H{
			"message": "添加flag失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})

	// 返回是否成功
}
