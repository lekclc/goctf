package team

import (
	team_ "src/logic/team"

	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Team) Info(c *gin.Context) {
	// token,team
	// mode, all or one
	var team Team
	c.ShouldBind(&team)
	var t team_.Team
	t.Name = team.Name
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
	members := strings.Split(msg["members"].(string), ",")
	var member []string
	member = append(member, members...)

	c.JSON(200, gin.H{
		"message":   "success",
		"name":      msg["name"],
		"leader":    msg["leader"],
		"members":   member,
		"desc":      msg["desc"],
		"score":     msg["score"],
		"challenge": challenge,
	})
}
