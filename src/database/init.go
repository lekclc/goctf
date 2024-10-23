package database

import (
	"fmt"
	"log"
	cfg "src/config"
	"strconv"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

type Db struct {
	db *gorm.DB
}

func NewDb() *Db {
	return &Db{}
}

func (db *Db) Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Cfg.Db.Uname, cfg.Cfg.Db.Passwd, cfg.Cfg.Db.Ip, strconv.Itoa(cfg.Cfg.Db.Port), "")
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	db.db = dbConn
	db.Set_db()
}

func (db *Db) Set_db() {
	//新建数据库
	db.db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.Cfg.Db.Dbname)
	//选择数据库
	db.db.Exec("USE " + cfg.Cfg.Db.Dbname)
	//创建表
	db.db.AutoMigrate(&User{})
	db.db.AutoMigrate(&Image{})
	db.db.AutoMigrate(&Challenge_ku{})
}
