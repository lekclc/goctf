package database

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	cfg "src/config"
	"src/ctrl/docker"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (db *Db) Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Cfg.Db.Uname, cfg.Cfg.Db.Passwd, cfg.Cfg.Db.Ip, strconv.Itoa(cfg.Cfg.Db.Port), "")
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to open database: %v\n", err)
		image := docker.NewImage("mysql")
		image.Get_info_by_name()
		if image.ID == "" {
			image.Port = map[int]int{cfg.Cfg.Db.Port: cfg.Cfg.Db.Port}
			image.Pull_image()
		}
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		cname := "mysql" + strconv.Itoa(time.Now().Nanosecond())
		cmd := exec.Command("docker", "run", "-d",
			"-p", "13306:3306",
			"-v", filepath.Join(dir, "/data/mysql/conf.d")+":/etc/mysql/conf.d",
			"-v", filepath.Join(dir, "data/mysql/log")+":/var/log/mysql",
			"-v", filepath.Join(dir, "data/mysql/data")+":/var/lib/mysql",
			"-e", "MYSQL_ROOT_PASSWORD="+cfg.Cfg.Db.Passwd,
			"-e", "MYSQL_ROOT_HOST=%",
			"--name", cname,
			"mysql")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			panic(fmt.Sprintf("Failed to run docker command: %v\n", err))
		}
		dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("Failed to open database: %v\n", err)
			os.Exit(1)
		}

	}

	db.Db = dbConn
	db.Set_db()
}
