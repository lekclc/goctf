package challenge_

import (
	"errors"
	con "src/const"
	"src/database"
)

func (s *Challenge) UpdateInfo() error {
	db := con.Db.Db
	var c database.Challenge
	db.Where("name = ?", s.Name).First(&c)
	if c.ID == 0 {
		return errors.New("challenge not found")
	}
	if s.Class != "" {
		c.Class = s.Class
	}
	if s.FileName != "" && s.FileName != c.FileName {
		c.FileName = s.FileName
	}
	if s.ImageName != "" && s.ImageName != c.ImageName {
		c.ImageName = s.ImageName
	}
	if s.Desc != "" && s.Desc != c.Desc {
		c.Desc = s.Desc
	}
	if s.MaxScore != 0 && s.MaxScore != c.MaxScore {
		c.MaxScore = s.MaxScore
	}
	err := db.Save(&c).Error
	if err != nil {
		return errors.New("update fail")
	}
	return nil

}
