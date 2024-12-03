package game

import (
	game_ "src/logic/game"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 队伍参加比赛
func (s *Game) Join(c *gin.Context) {
	gameid := c.PostForm("game")
	team := c.PostForm("team")

	id, err := strconv.ParseUint(gameid, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid game ID"})
		return
	}

	var g game_.Game
	g.ID = uint(id)
	if err := g.Join(team); err != nil {
		c.JSON(500, gin.H{"error": "Failed to join game"})
		return
	}
	c.JSON(200, gin.H{"message": "Success"})

}
