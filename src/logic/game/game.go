package game_

import "time"

type Game struct {
	Name  string
	ID    uint
	Start time.Time
	End   time.Time
	Desc  string
}

func GetGame() *Game {
	return &Game{}
}
