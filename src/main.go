package main

import (
	cfg "src/config"
	"src/minit"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	var s = gin.Default()
	minit.Init(s)
	s.Run(cfg.Cfg.Host.Ip + ":" + strconv.Itoa(cfg.Cfg.Host.Port))

}
