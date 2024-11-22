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

type Challenge struct {
	gorm.Model
	Active    bool   `gorm:"unique;not null"`
	ImageID   uint   `gorm:"not null"`
	ImageName string `gorm:"not null"`
	Flag      string `gorm:"unique"`
	File      string
}

type Container struct {
	gorm.Model
	ChallengeID uint   `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	Flag        string `gorm:"not null"`
}

type Game struct {
	gorm.Model
}
