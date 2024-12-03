package game_

import (
	"errors"
	con "src/const"
	"src/database"
)

func (s *Game) Show() (map[any]any, error) {
	db := con.Db.Db
	var g database.Game
	var res = make(map[any]any)
	db.Where("id = ?", s.ID).First(&g)
	if g.Name == "" || g.ID == 0 {
		err := errors.New("Game not found")
		return res, err
	}
	//关卡
	var challenges []database.Challenge
	db.Where("game_id = ?", g.ID).Find(&challenges)
	for _, c := range challenges {
		res[c.Name] = c.ID
	}
	return res, nil
}
