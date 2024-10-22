package router

import (
	"src/api/user"

	"github.com/gin-gonic/gin"
)

func (r *router) RouterUser(s *gin.Engine) {
	u := user.User{}
	s.GET("/login", u.Login)
}
