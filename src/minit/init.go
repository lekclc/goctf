package minit

import (
	cfg "src/config"
	"src/crtl/docker"
	"src/crtl/router"
	"src/db"

	"github.com/gin-gonic/gin"
)

func Init(s *gin.Engine) {
	cfg.Cfg.Init()
	docker_ := docker.NewDocker()
	docker_.Init()
	db_ := db.NewDb()
	db_.Init()
	router_ := router.NewRouter()
	router_.RouterInit(s)
}