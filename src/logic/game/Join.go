package game_

import (
	"errors"
	con "src/const"
	"src/database"
)

func (s *Game) Join(team string) error {
	db := con.Db.Db

	var g database.Game
	var t database.Team
	err := db.Where("id = ?", s.ID).First(&g).Error
	if err != nil {
		return err
	}
	err = db.Where("name = ?", team).First(&t).Error
	if err != nil {
		return err
	}
	if t.GameID != 0 {
		return errors.New("team already in game")
	}
	t.GameID = g.ID
	err = db.Save(&t).Error
	return err
}
