package challenge

import (
	challenge_ "src/logic/challenge"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 提交flag,进行检验
func (s *Challenge) FlagSubmit(c *gin.Context) {
	name := c.PostForm("name")
	flag := c.PostForm("flag")
	game_id := c.PostForm("game_id")
	gameid, _ := strconv.Atoi(game_id)
	team_id := c.PostForm("team_id")
	teamid, _ := strconv.Atoi(team_id)
	challenge_id := c.PostForm("challenge_id")
	challengeid, _ := strconv.Atoi(challenge_id)
	var ch challenge_.Challenge
	ch.Flags = flag
	err := ch.FlagSubmit(uint(gameid), uint(challengeid), uint(teamid), name)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
