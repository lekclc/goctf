package user_

import (
	con "src/const"
	"src/ctrl/utils"
	"src/database"
)

func (s *User) Register() (bool, error) {
	var user database.User
	db := con.Db.Db
	db.Where("name = ? ", s.Uname).First(&user)
	if user.ID != 0 {
		return false, nil
	}
	Passwd, _ := utils.Hash_passwd(s.Passwd)
	if Passwd == "" {
		return false, nil
	}
	user.Name = s.Uname
	user.Passwd = Passwd
	user.Admin = false
	user.Team = ""
	err := db.Create(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
