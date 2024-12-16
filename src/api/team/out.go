package team

import (
	team_ "src/logic/team"

	"github.com/gin-gonic/gin"
)

// 提出成员
func (s *Team) Out(c *gin.Context) {
	// token,name,team,who to out
	team := c.PostForm("team") // team name
	name := c.PostForm("name") // who done
	out := c.PostForm("out")   // who to out
	var t team_.Team
	t.Name = team
	t.Leader = name
	err := t.Out(out)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "out success",
	})

	// yes or no
	//if no, message
}
