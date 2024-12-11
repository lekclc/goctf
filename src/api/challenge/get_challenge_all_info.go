package challenge

import (
	"fmt"
	challenge_ "src/logic/challenge"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Challenge) GetChallengeAllInfo(c *gin.Context) {
	challenge_id := c.PostForm("challenge_id")
	challengeid, _ := strconv.Atoi(challenge_id)
	var ch challenge_.Challenge
	res, err := ch.GetChallengeAllInfo(uint(challengeid))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	rett := make(map[string]interface{})
	for key, value := range res {
		strKey := fmt.Sprintf("%v", key)
		rett[strKey] = value
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    rett,
	})
}
