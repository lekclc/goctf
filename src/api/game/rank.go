package game

import (
	game_ "src/logic/game"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Game) GetAllTeamRank(c *gin.Context) {
	game_id := c.PostForm("game_id")
	gameid, err := strconv.ParseUint(game_id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}
	var g game_.Game
	g.ID = uint(gameid)
	res, err := g.GetAllTeamRank()
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
