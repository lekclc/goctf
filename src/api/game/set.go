package game

import (
	game_ "src/logic/game"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 添加比赛
func (s *Game) Set(c *gin.Context) {
	//token,name,
	start_time_str := c.PostForm("start")
	end_time_str := c.PostForm("end")
	name := c.PostForm("game")
	desc := c.PostForm("desc")
	id_ := c.PostForm("id")
	parsedID, err := strconv.ParseUint(id_, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	const timeFormat = "2006-01-02T15:04"
	start_time, err := time.Parse(timeFormat, start_time_str)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid start time format"})
		return
	}
	end_time, err := time.Parse(timeFormat, end_time_str)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid end time format"})
		return
	}

	var g game_.Game

	g.Start = start_time
	g.End = end_time
	g.Name = name
	g.Desc = desc
	var id uint
	if uint(parsedID) != 0 {
		id, err = g.Set(uint(parsedID))
	} else {
		id, err = g.Set()
	}
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"id":      id,
	})
	//return success or fail
}
