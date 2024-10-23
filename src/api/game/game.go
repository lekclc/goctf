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
func (s *Game_router) GameList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "game_list",
	})
}
func (s *Game_router) GameDetail(c *gin.Context) {
}
func (s *Game_router) GameAdd(c *gin.Context) {
}
func (s *Game_router) GameDel(c *gin.Context) {
}
