package user_

import (
	con "src/const"
	"src/database"
)

func (s *User) Update() error {
	db := con.Db.Db
	var u database.User
	err := db.Where("name = ?", s.Uname).First(&u).Error
	if err != nil {
		return err
	}
	u.Passwd = s.Passwd
	err = db.Save(&u).Error
	if err != nil {
		return err
	}

	return nil
}
