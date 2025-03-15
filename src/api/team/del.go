package team

import (
	team_ "src/logic/team"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Team) Del(c *gin.Context) {
	name := c.PostForm("name")
	team_id := c.PostForm("team_id")
	teamIDUint64, err := strconv.ParseUint(team_id, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "invalid team_id",
		})
		return
	}
	var t team_.Team
	err = t.Del(name, uint(teamIDUint64))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})

}
