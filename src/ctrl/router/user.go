package router

import (
	user_router "src/api/user"
)

func (r *Router) RouterUser() {
	u := user_router.NewUser()
	g := r.s.Group("/user")
	g.POST("/login", u.Login)
	g.GET("/register", u.Register)
	g.GET("/logout", u.Logout)
}
