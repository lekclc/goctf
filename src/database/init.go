package database

import (
	"fmt"
	"log"
	"os"
	cfg "src/config"
	"src/ctrl/utils"
	"strconv"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

type Db struct {
	Db *gorm.DB
}

func NewDb() *Db {
	return &Db{}
}

func (db *Db) Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Cfg.Db.Uname, cfg.Cfg.Db.Passwd, cfg.Cfg.Db.Ip, strconv.Itoa(cfg.Cfg.Db.Port), "")
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	db.Db = dbConn
	db.Set_db()
}

func (db *Db) Set_db() {
	//新建数据库
	db.Db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.Cfg.Db.Dbname)
	//选择数据库
	db.Db.Exec("USE " + cfg.Cfg.Db.Dbname)
	//创建表
	db.Db.AutoMigrate(&User{})
	db.Db.AutoMigrate(&Image{})
	db.Db.AutoMigrate(&Challenge{})
	db.Db.AutoMigrate(&Game{})
	var user User
	user.Uname = cfg.Cfg.Admin.Uname
	user.Passwd, _ = utils.Hash_passwd(cfg.Cfg.Admin.Passwd)
	user.Admin = true
	if user.Passwd == "" {
		log.Fatalf("Failed to hash password")
		os.Exit(1)
	}
	var user_ User
	db.Db.First(&user_, "uname = ?", cfg.Cfg.Admin.Uname)
	if user_.ID == 0 {
		db.Db.Create(&user)
	} else {

		db.Db.Model(&user_).Update("passwd", user.Passwd)

	}

}
