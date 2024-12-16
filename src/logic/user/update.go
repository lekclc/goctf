package user_

import (
	"errors"
	con "src/const"
	"src/ctrl/utils"
	"src/database"
)

func (s *User) Update(new_passwd string) error {
	db := con.Db.Db
	var u database.User
	err := db.Where("name = ?", s.Uname).First(&u).Error
	if err != nil {
		return err
	}
	is := utils.Cmp_Passwd(u.Passwd, s.Passwd)

	if !is {
		return errors.New("passwd error")
	}

	u.Passwd, err = utils.Hash_passwd(new_passwd)
	if err != nil {
		return err
	}
	err = db.Save(&u).Error
	if err != nil {
		return err
	}
	return nil
}
