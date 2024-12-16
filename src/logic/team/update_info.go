package team_

import (
	"crypto/md5"
	"fmt"
	con "src/const"
	"src/database"
)

func (s *Team) UpdateInfo(iskey bool) error {
	var t database.Team
	db := con.Db.Db

	err := db.Where("id = ?", s.ID).First(&t).Error
	if err != nil {
		return err
	}
	t.Name = s.Name
	t.Desc = s.Desc
	if iskey {
		var new_key string
		new_key = fmt.Sprintf("%x", md5.Sum([]byte(new_key)))
		t.Key = new_key
	}
	err = db.Save(&t).Error
	if err != nil {
		return err
	}
	return nil
}
