package team

import (
	team_ "src/logic/team"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Team) UpdateInfo(c *gin.Context) {
	teamId := c.PostForm("team_id")
	new_name := c.PostForm("new_team_name")
	new_desc := c.PostForm("new_team_desc")
	iskey := c.PostForm("key")

	team_id, err := strconv.Atoi(teamId)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "参数错误",
		})
	}
	var t team_.Team

	t.ID = uint(team_id)
	t.Name = new_name
	t.Desc = new_desc
	is_key := false
	if iskey == "true" {
		is_key = true
	}
	err = t.UpdateInfo(is_key)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})

}
