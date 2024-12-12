package router

import (
	"src/api/game"
	"src/ctrl/mid"
)

func (r *Router) RouterGame() {
	u := game.NewGame()
	g := r.s.Group("/game")
	g.POST("/gamelist", u.GameList)
	g.POST("/getallteamrank", u.GetAllTeamRank)

	d := g.Group("/")
	d.Use(mid.AuthMid())
	d.POST("/join", u.Join)
	d.POST("/show", u.Show)
	d.POST("/getconinfo", u.GetConInfo)

	s := g.Group("/")
	s.Use(mid.AuthAdmin())
	s.POST("/set", u.Set)
	s.POST("/talk", u.Talk)
}
