package team

import (
	team_ "src/logic/team"

	"github.com/gin-gonic/gin"
)

func (s *Team) Out(c *gin.Context) {
	// token,name,team,who to out
	team := c.PostForm("team") // team name
	name := c.PostForm("name") // who done
	out := c.PostForm("out")   // who to out
	var t team_.Team
	t.Name = team
	t.Leader = name
	status := t.Out(out)

	if status == 0 {
		c.JSON(200, gin.H{
			"message": "out success",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "error",
		})
	}

	// yes or no
	//if no, message
}
