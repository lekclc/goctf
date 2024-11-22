package challenge

import "github.com/gin-gonic/gin"

type Challenge struct {
	Name      string
	MaxScore  int
	DoneNum   int
	Score     int
	FileName  string
	ImageName string
	Flags     []string
	Hints     []string
}

func NewChallenge() *Challenge {
	return &Challenge{}
}

func (s *Challenge) Challenge(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "challenge",
	})
}
