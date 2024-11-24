package user_

import (
	con "src/const"
	"src/ctrl/utils"
	"src/database"
)

func (s *User) Login() (bool, bool, error) {
	// TODO
	var user database.User
	db := con.Db.Db
	db.Where("name = ? ", s.Uname).First(&user)
	if user.ID == 0 {
		return false, false, nil
	}
	if !utils.Cmp_Passwd(user.Passwd, s.Passwd) {
		return false, false, nil
	}
	return true, user.Admin, nil
}
