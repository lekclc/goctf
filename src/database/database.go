package database

import (
	"gorm.io/gorm"
)

type Db struct {
	Db *gorm.DB
	Id string
}

func NewDb() *Db {
	return &Db{}
}
