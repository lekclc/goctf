package game

import (
	game_ "src/logic/game"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Game) GetConInfo(c *gin.Context) {
	team_id := c.PostForm("team_id")
	game_id := c.PostForm("game_id")
	name := c.PostForm("name")
	teamid, err := strconv.ParseUint(team_id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "team_id is invalid",
		})
	}
	gameid, err := strconv.ParseUint(game_id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "game_id is invalid",
		})
	}
	var g game_.Game
	res, err := g.GetConInfo(uint(gameid), uint(teamid), name)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    res,
	})
}
