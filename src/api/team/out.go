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
	} else if status == 1 {
		c.JSON(400, gin.H{
			"message": "team not found",
		})
	} else if status == 2 {
		c.JSON(400, gin.H{
			"message": "user not found",
		})
	} else if status == 3 {
		c.JSON(400, gin.H{
			"message": "user not in team",
		})
	} else if status == 4 {
		c.JSON(400, gin.H{
			"message": "team not in user",
		})
	} else if status == 5 {
		c.JSON(400, gin.H{
			"message": "error",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "error",
		})
	}

	// yes or no
	//if no, message
}
