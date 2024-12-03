package router

import (
	"src/api/game"
	"src/ctrl/mid"
)

func (r *Router) RouterGame() {
	u := game.NewGame()
	g := r.s.Group("/game")
	g.POST("/gamelist", u.GameList)

	d := g.Group("/")
	d.Use(mid.AuthMid())
	d.POST("/join", u.Join)
	d.POST("/show", u.Show)

	s := g.Group("/")
	s.Use(mid.AuthAdmin())
	s.POST("/set", u.Set)
	s.POST("/talk", u.Talk)
}
