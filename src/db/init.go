package db

import (
	"fmt"
	"src/crtl/docker"
)

type Db struct {
}

func NewDb() *Db {
	return &Db{}
}

func (db *Db) Init() {
	//利用docker拉取一个数据库服务
	//检测本地镜像
	mysql_ := docker.NewImage()
	mysql_.Get_info_by_name("mysql")
	if mysql_.ID == "" {
		mysql_.Pull_image("mysql")
	}
	fmt.Println(mysql_.ID)
	//启动一个容器
}
