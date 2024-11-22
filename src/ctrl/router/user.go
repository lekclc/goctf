package router

import (
	user "src/api/user"
)

func (r *Router) RouterUser() {
	u := user.NewUser()
	g := r.s.Group("/user")
	g.POST("/login", u.Login)
	g.POST("/register", u.Register)
	g.GET("/logout", u.Logout)
}
