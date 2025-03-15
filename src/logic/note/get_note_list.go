package note_

import (
	con "src/const"
	"src/database"
)

func (s *Note) GetNoteList(name string) (map[uint]string, error) {

	db := con.Db.Db
	tx := db.Begin()
	defer tx.Rollback()

	res := make(map[uint]string)

	if name == "" {
		var n []database.Note
		err := tx.Find(&n).Error
		if err != nil {
			return nil, err
		}
		for _, note := range n {
			res[note.ID] = note.Name
		}
	} else {
		var n []database.Note
		var u database.User
		err := tx.Where("name = ?", name).First(&u).Error
		if err != nil {
			return nil, err
		}
		err = tx.Where("user_id = ?", u.ID).Find(&n).Error
		if err != nil {
			return nil, err
		}
		for _, note := range n {
			res[note.ID] = note.Name
		}
	}
	return res, nil

}
