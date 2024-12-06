package challenge

import (
	challenge_ "src/logic/challenge"

	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Challenge) DelCon(c *gin.Context) {
	//获得动态环境
	challenge_id := c.PostForm("challenge_id")
	game_id := c.PostForm("game_id")
	team_id := c.PostForm("team_id")
	name := c.PostForm("name")
	challlengeid, err := strconv.ParseUint(challenge_id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid challenge_id"})
		return
	}
	teamid, err := strconv.ParseUint(team_id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid challenge_id"})
		return
	}
	gameid, err := strconv.ParseUint(game_id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid challenge_id"})
		return
	}
	challenge := challenge_.GetChallenge()

	err = challenge.DelCon(uint(challlengeid), uint(teamid), uint(gameid), name)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
