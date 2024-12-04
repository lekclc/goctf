package game

import (
	game_ "src/logic/game"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 队伍参加比赛
func (s *Game) Join(c *gin.Context) {
	gameid := c.PostForm("game_id")
	team := c.PostForm("team_id")
	name := c.PostForm("name")
	id, err := strconv.ParseUint(gameid, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid game ID"})
		return
	}

	var g game_.Game
	g.ID = uint(id)
	teamID, err := strconv.ParseUint(team, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid team ID"})
		return
	}

	if err := g.Join(uint(teamID), name); err != nil {
		c.JSON(500, gin.H{"error": "Failed to join game"})
		return
	}
	c.JSON(200, gin.H{"message": "Success"})

}
