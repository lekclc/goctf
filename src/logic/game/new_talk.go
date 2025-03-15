package game_

import (
	con "src/const"
	"src/database"
)

func (s *Game) NewTalk(talk string, gameID int) (int, error) {
	db := con.Db.Db
	var t database.Talk
	tx := db.Begin()
	defer tx.Rollback()
	t.GameID = uint(gameID)
	t.Talk = talk
	err := tx.Create(&t).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit().Error
	if err != nil {
		return 0, err
	}
	return int(t.ID), nil
}
