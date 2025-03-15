package router

import (
	"src/api/team"
	"src/ctrl/mid"
)

func (r *Router) RouterTeam() {
	u := team.NewTeam()
	g := r.s.Group("/team")

	d := g.Group("/")
	d.Use(mid.AuthMid())
	d.POST("/create", u.Create)
	d.POST("/out", u.Out)
	d.POST("/join", u.Join)
	d.POST("/info", u.Info)
	g.POST("/updateinfo", u.UpdateInfo)
	g.POST("/del", u.Del)
}
