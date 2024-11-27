package challenge_

import (
	"errors"
	con "src/const"
	"src/database"
	"strings"
)

func (s *Challenge) FlagSubmit(name string) error {
	db := con.Db.Db
	var c database.Challenge
	db.Where("name = ?", s.Name).First(&c)
	if c.ID == 0 {
		return errors.New("challenge not found")
	}
	if !c.Active {
		flags := strings.Split(c.Flags, ",")
		for _, f := range flags {
			if f == s.Flags {
				return nil
			}
		}
	} else {
		var user database.User
		var Container database.Container
		db.Where("name = ?", name).First(&user)
		if user.ID == 0 {
			return errors.New("user not found")
		}
		db.Where("user_id = ? and challenge_id = ?", user.ID, c.ID).First(&Container)
		if Container.ID == 0 {
			return errors.New("container not found")
		}
		if Container.Flag != "" && Container.Flag == s.Flags {
			return nil
		}
	}
	return errors.New("flag error")
}
