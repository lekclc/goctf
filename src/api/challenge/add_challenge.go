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
	if dc.Class == "" || dc.Class != "PWN" && dc.Class != "REVERSE" && dc.Class != "WEB" && dc.Class != "MISC" && dc.Class != "CRYPTO" {
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
	port := c.PostForm("port")
	if port != "" {
		dc.Port = port
	}

	dc.MaxScore = uint(maxScore)
	dc.ImageID = 0
	dc.DoneNum = 0
	dc.Score = 0
	dc.ImageName = c.PostForm("image_name")
	dc.Flags = c.PostForm("flags")
	dc.Desc = c.PostForm("desc")

	gameidStr := c.PostForm("game_id")
	gameid, err := strconv.ParseUint(gameidStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid game_id",
		})
		return
	}

	dc.GameID = uint(gameid)
	status := dc.AddChallenge()
	if status != nil {
		c.JSON(400, gin.H{
			"message": status,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
