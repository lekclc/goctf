package router

import (
	game_router "src/api/game"
	"src/ctrl/mid"
)

func (r *Router) RouterGame() {
	game := game_router.NewGame()
	g := r.s.Group("/game")
	g.Use(mid.AuthMid())
	g.POST("/game", game.Game)
}
