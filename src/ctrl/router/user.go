package router

import (
	user "src/api/user"
	"src/ctrl/mid"
)

func (r *Router) RouterUser() {
	u := user.NewUser()
	g := r.s.Group("/user")
	g.POST("/login", u.Login)
	g.POST("/register", u.Register)

	s := g.Group("/")
	s.Use(mid.AuthMid())
	s.POST("/logout", u.Logout)
	s.POST("/info", u.Info)
}
