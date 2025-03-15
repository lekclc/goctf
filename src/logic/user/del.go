package user_

import (
	"errors"
	con "src/const"
	"src/database"
)

func (s *User) Del(name string) error {
	db := con.Db.Db
	var u database.User
	u.Name = name
	err := db.First(&u).Error
	if err != nil {
		return errors.New("no user")
	}
	err = db.Delete(&u).Error
	//其余相关数据操作由数据库中的触发器负责
	if err != nil {
		return err
	}

	return nil
}
