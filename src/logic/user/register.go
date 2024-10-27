package user_

import (
	con "src/const"
	"src/ctrl/utils"
	"src/database"
)

func (s *User_) Register() (bool, error) {
	var user database.User
	db := con.Db.Db
	db.Where("uname = ? ", s.Uname).First(&user)
	if user.ID != 0 {
		return false, nil
	}
	Passwd, _ := utils.Hash_passwd(s.Passwd)
	if Passwd == "" {
		return false, nil
	}
	user.Uname = s.Uname
	user.Passwd = Passwd
	user.Admin = false
	db.Create(&user)
	return true, nil
}
