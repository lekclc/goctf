package con

import (
	"src/ctrl/docker"
	"src/database"
)

// 一些方便全局使用的变量

// docker连接, 用于操作docker
/*
	ctx context.Context
	cl  *client.Client
*/
var Docker *docker.Docker

// 数据库连接
/*
	Db *gorm.DB
	Id string
*/
var Db *database.Db
