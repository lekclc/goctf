package game_

import (
	"errors"
	con "src/const"
	"src/database"
)

func (s *Game) Set() (uint, error) {

	db := con.Db.Db
	var g database.Game
	err := db.Where("name = ?", s.Name).First(&g).Error
	if err != nil {
		return 0, err
	}
	if g.ID != 0 {
		return 0, errors.New("Game already exists")
	}
	g.Name = s.Name
	g.Start = s.Start
	g.End = s.End
	err = db.Create(&g).Error
	if err != nil {
		return 0, err
	}
	return g.ID, nil
}