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
	err := t.Join(name)

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
