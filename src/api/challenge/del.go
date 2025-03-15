package challenge

import (
	challenge_ "src/logic/challenge"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Challenge) Del(c *gin.Context) {
	challenge_id := c.PostForm("challenge_id")
	teamIDUint64, err := strconv.ParseUint(challenge_id, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "invalid challenge_id",
		})
		return
	}
	var t challenge_.Challenge
	err = t.Del(uint(teamIDUint64))
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
