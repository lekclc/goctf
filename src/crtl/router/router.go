package router

import "github.com/gin-gonic/gin"

type router struct{}

func NewRouter() *router {
	return &router{}
}

func (r *router) RouterInit(s *gin.Engine) {
	r.RouterUser(s)
}
