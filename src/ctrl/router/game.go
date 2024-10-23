package router

import (
	game_router "src/api/game"
)

func (r *Router) RouterGame() {
	game := game_router.NewGame()
	g := r.s.Group("/game")
	g.GET("/game", game.Game)
}
