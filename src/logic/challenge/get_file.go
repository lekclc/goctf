package challenge_

import (
	"errors"
	"path/filepath"
	con "src/const"
	"src/database"
)

func (s *Challenge) GetFile() (string, error) {

	var dc database.Challenge
	db := con.Db.Db

	db.Where("name = ?", s.Name).First(&dc)
	if dc.ID == 0 {
		return "", errors.New("challenge not found")
	}
	if dc.FileName == "" {
		return "", errors.New("no file")
	}
	filePath := filepath.Join("data/file", dc.FileName)
	return filePath, nil

}
