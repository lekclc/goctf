package game

import (
	game_ "src/logic/game"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Game) Set(c *gin.Context) {
	//token,name,
	start_time_str := c.PostForm("start")
	end_time_str := c.PostForm("end")
	name := c.PostForm("name")

	start_time, err := time.Parse(time.RFC3339, start_time_str)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid start time format"})
		return
	}
	end_time, err := time.Parse(time.RFC3339, end_time_str)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid end time format"})
		return
	}

	var g game_.Game

	g.Start = start_time
	g.End = end_time
	g.Name = name
	id, err := g.Set()
	if id == 0 || err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"id":      id,
	})
	//return success or fail
}
