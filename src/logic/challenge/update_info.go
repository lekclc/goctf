package challenge_

import (
	"errors"
	con "src/const"
	"src/database"
)

func (s *Challenge) UpdateInfo(challengeid uint) error {
	db := con.Db.Db
	var c database.Challenge
	err := db.Where("id = ?", challengeid).First(&c).Error
	if err != nil {
		return errors.New("challenge not found")
	}
	if !c.Active {
		if s.Flags != "" && s.Flags != c.Flags {
			c.Flags = s.Flags
		}
	} else if c.Active {
		if s.ImageName != "" && s.ImageName != c.ImageName {
			c.ImageName = s.ImageName
		}
	}
	if s.Port != "" && s.Port != c.Port {
		c.Port = s.Port
	}
	if s.Desc != "" && s.Desc != c.Desc {
		c.Desc = s.Desc
	}
	if s.MaxScore != 0 && s.MaxScore != c.MaxScore {
		c.MaxScore = s.MaxScore
	}
	err = db.Save(&c).Error
	if err != nil {
		return errors.New("update fail")
	}
	return nil

}
