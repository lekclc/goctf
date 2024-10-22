package main

import (
	router "src/crtl/router"
	"src/minit"

	"github.com/gin-gonic/gin"
)

func main() {

	var s = gin.Default()
	minit.Init()
	router := router.NewRouter()
	router.RouterInit(s)
	s.Run(":8000")

}
