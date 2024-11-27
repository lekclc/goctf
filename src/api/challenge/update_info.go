package challenge

import (
	challenge_ "src/logic/challenge"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加hint
func (s *Challenge) UpdateInfo(c *gin.Context) {
	var dc challenge_.Challenge
	dc.Name = c.PostForm("challenge_name") //无法更改
	dc.FileName = c.PostForm("file_name")
	dc.ImageName = c.PostForm("image_name")
	dc.Desc = c.PostForm("desc")
	maxscore := c.PostForm("max_score")
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
	err := dc.UpdateInfo()
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
