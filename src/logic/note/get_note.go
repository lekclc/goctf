package note_

import (
	"encoding/base64"
	"os"
	con "src/const"
	"src/database"
)

func (s *Note) GetNote(id uint) (string, error) {
	db := con.Db.Db
	tx := db.Begin()
	defer tx.Rollback()

	var n database.Note
	if err := tx.Where("id = ?", id).First(&n).Error; err != nil {
		return "", err
	}
	data, err := os.ReadFile(n.Path)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, nil
}
