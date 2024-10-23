package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uname  string `gorm:"unique;not null;size:255"`
	Passwd string `gorm:"unique;not null;size:255"`
}

type Image struct {
	Id int `json:"id"`
}

type Challenge_ku struct {
	Id int `json:"id"`
}
