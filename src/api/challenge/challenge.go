package challenge

import "github.com/gin-gonic/gin"

type Challenge struct{}

func NewChallenge() *Challenge {
	return &Challenge{}
}

func (s *Challenge) Challenge(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "challenge",
	})
}
