package challenge

import (
	challenge_ "src/logic/challenge"

	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Challenge) GetCon(c *gin.Context) {
	//token, 题目名称, name
	//获得动态环境
	name := c.PostForm("name")
	challenge_id := c.PostForm("challenge_id")
	id, err := strconv.ParseUint(challenge_id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid challenge_id"})
		return
	}
	challenge := challenge_.GetChallenge()

	res, err := challenge.GetCon(uint(id), name)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    res,
	})

	//return host,port
}
