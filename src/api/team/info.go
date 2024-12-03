package team

import (
	team_ "src/logic/team"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取队伍信息
func (s *Team) Info(c *gin.Context) {
	// token,team
	team_id := c.PostForm("team_id")
	id, err := strconv.ParseUint(team_id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid team_id",
		})
		return
	}
	var t team_.Team
	t.ID = uint(id)
	msg, err := t.Info()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}
	challenges := strings.Split(msg["challenge"].(string), ",")
	var challenge []string
	challenge = append(challenge, challenges...)
	challenge = challenge[:len(challenge)-1]
	members := strings.Split(msg["members"].(string), ",")
	var member []string
	member = append(member, members...)
	member = member[:len(member)-1]
	c.JSON(200, gin.H{
		"message":   "success",
		"name":      msg["name"],
		"leader":    msg["leader"],
		"members":   member,
		"desc":      msg["desc"],
		"score":     msg["score"],
		"challenge": challenge,
		"gameID":    msg["gameid"],
		"key":       msg["key"],
	})
}
