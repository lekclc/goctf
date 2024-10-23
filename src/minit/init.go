package minit

import (
	cfg "src/config"
	con "src/const"
	"src/ctrl/docker"
	"src/ctrl/router"
	database "src/database"

	"github.com/gin-gonic/gin"
)

func Init(s *gin.Engine) {
	cfg.Cfg.Init()
	con.Docker = docker.NewDocker()
	con.Docker.Init()
	con.Db = database.NewDb()
	con.Db.Init()
	router_ := router.NewRouter(s)
	router_.RouterInit()
}
