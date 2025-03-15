package game

import (
	game_ "src/logic/game"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Game) NewTalk(c *gin.Context) {
	// token,name,消息内容,比赛id或者0表示主界面
	talk := c.PostForm("talk")
	game_id := c.PostForm("game_id")
	gameID, err := strconv.ParseInt(game_id, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "game_id error",
		})
		return
	}
	var g game_.Game
	talk_id, err := g.NewTalk(talk, int(gameID))
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"talk_id": talk_id,
	})
}
