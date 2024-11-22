package game

import "time"

type Game struct {
	Name  string
	Id    uint
	Start time.Time
	End   time.Time
}

func NewGame() *Game {
	return &Game{}
}
