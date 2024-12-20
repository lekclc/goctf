package challenge_

import (
	con "src/const"
	"src/database"
)

func (s *Challenge) AddChallenge() int {

	db := con.Db.Db
	var dc database.Challenge
	db.Where("name = ?", s.Name).First(&dc)
	if dc.ID != 0 {
		return 1
		// return challenge exist
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
	if dc.ImageName != "" {
		db.Where("name = ?", dc.ImageName).First(&image)
		if image.ID != 0 {
			return 3
		}
		image.Name = dc.ImageName
		image.Port = dc.Port
		db.Create(&image)
		dc.ImageID = image.ID
	}
	err := db.Create(&dc).Error
	if err != nil {
		return 2
		// return fail
	}
	return 0
	//return success or fail

}
