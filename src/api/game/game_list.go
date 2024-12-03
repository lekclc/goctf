package game

import (
	"fmt"
	game_ "src/logic/game"

	"github.com/gin-gonic/gin"
)

// 获取比赛列表
func (s *Game) GameList(c *gin.Context) {
	var g game_.Game
	res, err := g.GameList()
	fmt.Println(res, err)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    res,
	})
}
