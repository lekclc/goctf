package router

import "github.com/gin-gonic/gin"

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) RouterInit(s *gin.Engine) {
	r.RouterUser(s)
}
