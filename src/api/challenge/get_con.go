package challenge

import (
	cfg "src/config"
	challenge_ "src/logic/challenge"

	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Challenge) GetCon(c *gin.Context) {
	//获得动态环境
	challenge_id := c.PostForm("challenge_id")
	game_id := c.PostForm("game_id")
	team_id := c.PostForm("team_id")
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

	res, err := challenge.GetCon(uint(challlengeid), uint(teamid), uint(gameid))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"host":    cfg.Cfg.Host.Ip,
		"port":    res,
	})
}
