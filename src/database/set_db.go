package database

import (
	"log"
	"os"
	cfg "src/config"
	"src/ctrl/utils"
)

func (db *Db) Set_db() {

	//创建表
	db.Db.AutoMigrate(&User{})
	db.Db.AutoMigrate(&Team{})
	db.Db.AutoMigrate(&Challenge{})
	db.Db.AutoMigrate(&Image{})
	db.Db.AutoMigrate(&Container{})
	db.Db.AutoMigrate(&Game{})
	db.Db.AutoMigrate(&Talk{})
	db.Db.AutoMigrate(&Note{})
	var user User
	user.Name = cfg.Cfg.Admin.Uname
	user.Passwd, _ = utils.Hash_passwd(cfg.Cfg.Admin.Passwd)
	user.Admin = true
	if user.Passwd == "" {
		log.Fatalf("Failed to hash password")
		os.Exit(1)
	}
	var user_ User
	db.Db.First(&user_, "name = ?", cfg.Cfg.Admin.Uname)
	if user_.ID == 0 {
		db.Db.Create(&user)
	} else {
		db.Db.Model(&user_).Update("passwd", user.Passwd)

	}

}
