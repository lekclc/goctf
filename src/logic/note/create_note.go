package note_

import (
	"crypto/sha256"
	"errors"
	"fmt"
	con "src/const"
	"src/database"
	"time"
)

func (s *Note) CreateNote(name string, note_name string) (string, error) {
	db := con.Db.Db
	var n database.Note
	var u database.User
	tx := db.Begin()
	defer tx.Rollback()
	err := tx.First(&u, "name = ?", name).Error
	if err != nil {
		return "", errors.New("user not found")
	}
	tx.First(&n, "user_id = ? and name = ?", u.ID, note_name)
	if n.ID != 0 {
		return "", errors.New("note already exists")
	}
	hash := sha256.Sum256([]byte(name + time.Now().String()))
	path := "data/note/" + fmt.Sprintf("%x", hash)

	n.Name = note_name
	n.UserId = u.ID
	n.Path = path
	err = tx.Create(&n).Error
	if err != nil {
		return "", err
	}
	err = tx.Commit().Error
	if err != nil {
		return "", err
	}
	return path, nil
}
