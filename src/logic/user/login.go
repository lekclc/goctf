package user_

import (
	con "src/const"
	"src/ctrl/utils"
	"src/database"
)

func (s *User_) Login() (bool, error) {
	// TODO
	var user database.User
	db := con.Db.Db
	db.Where("uname = ? ", s.Uname).First(&user)
	if user.ID == 0 {
		return false, nil
	}
	if !utils.Cmp_Passwd(user.Passwd, s.Passwd) {
		return false, nil
	}
	return true, nil
}
