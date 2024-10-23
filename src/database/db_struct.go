package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uname  string `gorm:"unique;not null;size:255"`
	Passwd string `gorm:"unique;not null;size:255"`
	Admin  bool   `gorm:"not null"`
}

type Image struct {
	gorm.Model
}

type Challenge struct {
	gorm.Model
	Active  bool   `gorm:"unique;not null"`
	ImageID uint   `gorm:"not null"`
	Flag    string `gorm:"unique"`
}

type Game struct {
	gorm.Model
}
