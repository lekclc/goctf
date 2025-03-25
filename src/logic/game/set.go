package game_

import (
	"errors"
	con "src/const"
	"src/database"
)

func (s *Game) Set(id ...uint) (uint, error) {

	db := con.Db.Db
	var g database.Game
	db.Where("name = ?", s.Name).First(&g)
	g.Name = s.Name
	g.Start = s.Start
	g.End = s.End
	g.Desc = s.Desc
	var err error
	if id != nil {
		if g.ID != id[0] {
			return 0, errors.New("Game not exists")
		}
		err = db.Save(&g).Error
	} else {
		if g.ID != 0 {
			return 0, errors.New("Game already exists")
		}
		err = db.Create(&g).Error
	}

	if err != nil {
		return 0, err
	}
	return g.ID, nil
}
