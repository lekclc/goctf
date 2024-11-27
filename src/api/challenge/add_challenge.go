package challenge

import (
	challenge_ "src/logic/challenge"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 赛题添加
func (s *Challenge) AddChallenge(c *gin.Context) {
	// class, mode (active or static), name, score,...description, ...flag, ...file, ...con
	var dc challenge_.Challenge
	dc.Class = c.PostForm("class")
	if dc.Class == "" || dc.Class != "PWN" && dc.Class != "RE" && dc.Class != "WEB" && dc.Class != "MISC" && dc.Class != "CRYPTO" {
		c.JSON(400, gin.H{
			"message": "class is empty or error",
		})
		return
	}
	dc.Name = c.PostForm("challenge_name")
	dc.Active = c.PostForm("active") == "true"
	maxScore, err := strconv.ParseUint(c.PostForm("max_score"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid max_score",
		})
		return
	}
	dc.MaxScore = uint(maxScore)
	dc.ImageID = 0
	dc.DoneNum = 0
	dc.Score = 0
	dc.FileName = c.PostForm("file_name")
	dc.ImageName = c.PostForm("image_name")
	dc.Flags = c.PostForm("flags")
	dc.Desc = c.PostForm("desc")
	status := dc.AddChallenge()
	if status == 1 {
		c.JSON(400, gin.H{
			"message": "challenge exist",
		})
	} else if status == 2 {
		c.JSON(400, gin.H{
			"message": "error",
		})
	} else if status == 0 {
		c.JSON(200, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "error",
		})
	}
}
