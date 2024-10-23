package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	s *gin.Engine
}

func NewRouter(s *gin.Engine) *Router {
	return &Router{s}
}

func (r *Router) RouterInit() {
	r.RouterUser()
	r.RouterGame()
	r.RouterChallenge()
}
