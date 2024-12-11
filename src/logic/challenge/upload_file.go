package challenge_

import (
	"os"
	"path"
	con "src/const"
	"src/database"
)

func (s *Challenge) UploadFile(challengeid uint, filename string) error {
	db := con.Db.Db
	var c database.Challenge
	err := db.Where("id = ?", challengeid).First(&c).Error
	if err != nil {
		return err
	}
	if c.FileName != "" {
		path := path.Join("data/file", c.FileName)
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	c.FileName = filename
	err = db.Save(&c).Error
	if err != nil {
		return err
	}
	return nil
}
