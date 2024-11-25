package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `gorm:"unique;not null;size:64"`
	Passwd  string `gorm:"unique;not null;size:255"`
	Admin   bool   `gorm:"unique;not null"`
	Team    string `gorm:"size:512"`
	Game    string `gorm:"size:512"`
	TeamNum uint   `gorm:"not null"` //最大队伍数为15
}

type Team struct {
	gorm.Model
	Name      string `gorm:"unique;not null;size:64"`
	Leader    string `gorm:"unique;not null;size:64"`
	Member    string `gorm:"unique;not null;size:255"`
	MemberNum uint   `gorm:"not null"` //最大队伍人数为4
	Desc      string `gorm:"size:255"`
	Key       string `gorm:"unique;not null;size:64"`
	Challenge string `gorm:"size:1024"`
	Score     uint   `gorm:"not null"`
	GameID    uint   `gorm:"not null"`
}

type Challenge struct {
	gorm.Model
	Active    bool   `gorm:"unique;not null"`
	Name      string `gorm:"unique;not null;size:64"`
	MaxScore  int    `gorm:"not null"`
	ImageID   uint   `gorm:"not null"`
	DoneNum   uint   `gorm:"not null"`
	Score     uint   `gorm:"not null"`
	FileName  string `gorm:"not null"`
	ImageName string `gorm:"not null"`
	Flags     string `gorm:"unique;size:512"`
	Hints     string `gorm:"size:512"`
}

type Image struct {
	gorm.Model
	Name        string `gorm:"unique;not null;size:64"`
	Port        string `gorm:"unique;not null;size:256"`
	ChallengeID uint   `gorm:"not null"`
}

type Container struct {
	gorm.Model
	Port        string `gorm:"unique;not null;size:256"`
	Flag        string `gorm:"unique;not null"`
	ChallengeID uint   `gorm:"unique;not null"`
	UserID      uint   `gorm:"unique;not null"`
}

type Game struct {
	gorm.Model
	Name  string `gorm:"unique;not null;size:64"`
	Start time.Time
	End   time.Time
}
