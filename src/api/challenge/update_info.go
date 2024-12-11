package challenge

import (
	challenge_ "src/logic/challenge"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加hint
func (s *Challenge) UpdateInfo(c *gin.Context) {
	var dc challenge_.Challenge
	challenge_id := c.PostForm("challenge_id")
	challengeid, _ := strconv.Atoi(challenge_id)
	dc.ImageName = c.PostForm("image_name")
	dc.Flags = c.PostForm("flags")
	dc.Desc = c.PostForm("desc")
	maxscore := c.PostForm("max_score")
	dc.Port = c.PostForm("port")
	if maxscore != "" {
		maxScore, err := strconv.ParseUint(maxscore, 10, 32)
		dc.MaxScore = uint(maxScore)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "error",
			})
			return
		}
	}

	err := dc.UpdateInfo(uint(challengeid))
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
