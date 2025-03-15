package game_

import (
	con "src/const"
	"src/database"
)

func (s *Game) GetTalk(gameID uint) (map[int]string, error) {
	db := con.Db.Db
	tx := db.Begin()
	defer tx.Rollback()
	var talks []database.Talk
	err := tx.Where("game_id = ?", gameID).Find(&talks).Error
	if err != nil {
		return nil, err
	}
	res := make(map[int]string)
	for _, talk := range talks {
		res[int(talk.ID)] = talk.Talk
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
