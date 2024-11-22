package game_router

import "github.com/gin-gonic/gin"

type Game_router struct {
}

func NewGame() *Game_router {
	return &Game_router{}
}
func (s *Game_router) Game(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "game",
	})
}
