package challenge_

import (
	con "src/const"
	"src/database"
)

func (s *Challenge) Del(challengeID uint) error {
	db := con.Db.Db
	var c database.Challenge

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := tx.First(&c, challengeID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Delete(&c).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
