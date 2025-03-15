package game

import (
	game_ "src/logic/game"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Game) GetTalk(c *gin.Context) {
	game_id := c.PostForm("game_id")
	// 0 作为主页公告
	gameID, err := strconv.ParseUint(game_id, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "game_id error",
		})
		return
	}
	var g game_.Game
	res, err := g.GetTalk(uint(gameID))
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "logic error",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    res,
	})
}
