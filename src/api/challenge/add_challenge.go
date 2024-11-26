package challenge

import (
	"fmt"
	challenge_ "src/logic/challenge"

	"github.com/gin-gonic/gin"
)

// 赛题添加
func (s *Challenge) AddChallenge(c *gin.Context) {
	// class, mode (active or static), name, score,...description, ...flag, ...file, ...con

	var info Challenge
	c.ShouldBind(&info)
	fmt.Println(info)
	var dc challenge_.Challenge
	dc.Class = info.Class
	dc.Name = info.Name
	dc.Active = info.Active
	dc.MaxScore = info.MaxScore
	dc.ImageID = info.ImageID
	dc.DoneNum = info.DoneNum
	dc.Score = info.Score
	dc.FileName = info.FileName
	dc.ImageName = info.ImageName
	dc.Flags = info.Flags
	dc.Desc = info.Desc
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
