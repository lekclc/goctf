package team

import (
	team_ "src/logic/team"

	"github.com/gin-gonic/gin"
)

// 加入队伍
func (s *Team) Join(c *gin.Context) {
	// token,name,team,key
	team_name := c.PostForm("team")
	name := c.PostForm("name")
	key := c.PostForm("key") //ShouldBind绑定失败
	var t team_.Team
	t.Name = team_name
	t.Key = key
	status := t.Join(name)

	if status == 1 {
		c.JSON(400, gin.H{
			"message": "team full",
		})
	} else if status == 2 {
		c.JSON(400, gin.H{
			"message": "key error",
		})
	} else if status == 0 {
		c.JSON(200, gin.H{
			"message": "success",
		})
	} else if status == 3 {
		c.JSON(400, gin.H{
			"message": "user not exist",
		})
	} else if status == 4 {
		c.JSON(400, gin.H{
			"message": "user already in team",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "error",
		})
	}
}
