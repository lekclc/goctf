package router

import (
	"src/api/user"

	"github.com/gin-gonic/gin"
)

func (r *Router) RouterUser(s *gin.Engine) {
	u := user.NewUser()
	g := s.Group("/user")
	g.GET("/login", u.Login)
	g.GET("/register", u.Register)
	g.GET("/logout", u.Logout)

}
