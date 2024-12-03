package game

import (
	game_ "src/logic/game"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取相关比赛的题目信息
func (s *Game) Show(c *gin.Context) {
	// token,name
	game_id := c.PostForm("game_id")
	var g game_.Game
	id, err := strconv.ParseUint(game_id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid game_id"})
		return
	}
	g.ID = uint(id)
	res, err := g.Show()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    res,
	})
}
