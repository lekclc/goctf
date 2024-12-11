package router

import (
	"src/api/challenge"
	"src/ctrl/mid"
)

func (r *Router) RouterChallenge() {
	u := challenge.NewChallenge()
	g := r.s.Group("/challenge")

	d := g.Group("/")
	d.Use(mid.AuthMid())
	d.POST("/flagsubmit", u.FlagSubmit)
	d.POST("/getcon", u.GetCon)
	d.POST("/delcon", u.DelCon)
	d.POST("/getfile", u.GetFile)
	d.POST("/info", u.Info)
	d.POST("/start", u.ChallengeStart)

	s := g.Group("/")
	s.Use(mid.AuthAdmin())
	s.POST("/addchallenge", u.AddChallenge)
	s.POST("/addflag", u.AddFlag)
	s.POST("/updateinfo", u.UpdateInfo)
	s.POST("/updatefile", u.UpdateFile)
	s.POST("/getallchallengeinfo", u.GetChallengeAllInfo)

}
