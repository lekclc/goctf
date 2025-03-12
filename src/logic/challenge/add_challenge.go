package challenge_

import (
	"fmt"
	con "src/const"
	"src/database"
)

func (s *Challenge) AddChallenge() error {

	db := con.Db.Db
	var dc database.Challenge
	err := db.Where("name = ?", s.Name).First(&dc).Error
	if dc.ID != 0 {
		return err
	}

	dc.Class = s.Class
	dc.Name = s.Name
	dc.Active = s.Active
	dc.MaxScore = s.MaxScore
	dc.ImageID = s.ImageID
	dc.DoneNum = s.DoneNum
	dc.Score = s.Score
	dc.FileName = s.FileName
	dc.ImageName = s.ImageName
	dc.GameID = s.GameID
	dc.Score = dc.MaxScore
	dc.Port = s.Port
	if s.Flags != "" {
		dc.Flags = s.Flags + ","
	}
	dc.Desc = s.Desc
	var image database.Image

	if dc.Active {
		if dc.ImageName != "" {
			err = db.Where("name = ?", dc.ImageName).First(&image).Error
			if image.ID != 0 {
				return err
			}
			image.Name = dc.ImageName
			image.Port = dc.Port
			db.Create(&image)
			dc.ImageID = image.ID
		}
	}
	err = db.Create(&dc).Error
	if err != nil {
		return err
		// return fail
	}
	fmt.Println("success")
	return nil
	//return success

}
