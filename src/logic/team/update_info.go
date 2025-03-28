package team_

import (
	"crypto/md5"
	"fmt"
	con "src/const"
	"src/database"
	"time"
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
		new_key := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()+t.Name+t.Key)))
		t.Key = new_key
	}
	err = db.Save(&t).Error
	if err != nil {
		return err
	}
	return nil
}
