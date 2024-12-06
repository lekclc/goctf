package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `gorm:"unique;not null;size:64"`
	Passwd  string `gorm:"not null;size:255"`
	Admin   bool   `gorm:"not null"`
	Team    string `gorm:"size:512"`
	TeamNum uint   `gorm:"not null"` //最大队伍数为15
}

type Team struct {
	gorm.Model
	Name      string `gorm:"unique;not null;size:64"`
	Leader    string `gorm:"not null;size:64"`
	Member    string `gorm:"not null;size:255"`
	MemberNum uint   `gorm:"not null"` //最大队伍人数为4
	Desc      string `gorm:"size:255"`
	Key       string `gorm:"not null;size:64"`
	Challenge string `gorm:"size:1024"`
	Score     uint   `gorm:"not null"`
	GameID    uint   `gorm:"not null"`
}

type Challenge struct {
	gorm.Model
	Class     string `gorm:"not null"`
	Active    bool   `gorm:"not null"`
	Name      string `gorm:"unique;not null;size:64"`
	MaxScore  uint   `gorm:"not null"`
	ImageID   uint   `gorm:"not null"`
	DoneNum   uint   `gorm:"not null"`
	Score     uint   `gorm:"not null"`
	FileName  string `gorm:"size:64"`
	ImageName string `gorm:"size:64"`
	Flags     string `gorm:"unique;size:512"`
	Desc      string `gorm:"size:512"`
	GameID    uint   `gorm:"not null"`
	Port      string `gorm:"size:64"`
}

type Image struct {
	gorm.Model
	Name        string `gorm:"unique;not null;size:64"`
	Port        string `gorm:"unique;not null;size:256"`
	ChallengeID uint   `gorm:"not null"`
}

type Container struct {
	gorm.Model
	ContainerID string `gorm:"unique;not null;size:64"`
	Port        string `gorm:"unique;not null;size:256"`
	Flag        string `gorm:"not null;size:128"`
	ChallengeID uint   `gorm:"not null"`
	TeamID      uint   `gorm:"not null"`
	GameID      uint   `gorm:"not null"`
}

type Game struct {
	gorm.Model
	Name  string `gorm:"unique;not null;size:64"`
	Start time.Time
	End   time.Time
	Desc  string `gorm:"size:512"`
}
