package router

import "src/ctrl/mid"

func (r *Router) RouterInit() {
	r.s.Use(mid.Cors())
	r.s.Static("/file", "./data/file/")
	r.RouterUser()
	r.RouterTeam()
	r.RouterGame()
	r.RouterChallenge()
}
