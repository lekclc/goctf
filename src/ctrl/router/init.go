package router

import (
	"src/ctrl/mid"

	"github.com/gin-gonic/gin"
)

type Router struct {
	s *gin.Engine
}

func NewRouter(s *gin.Engine) *Router {
	return &Router{s}
}

func (r *Router) RouterInit() {
	r.s.Use(mid.Cors())
	r.RouterUser()
	r.RouterGame()
	r.RouterChallenge()
}
