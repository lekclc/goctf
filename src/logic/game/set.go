package game_

import (
	"errors"
	con "src/const"
	"src/database"
	"time"
)

func (s *Game) Set(id ...uint) (uint, error) {

	db := con.Db.Db
	var g database.Game
	db.Where("name = ?", s.Name).First(&g)
	g.Name = s.Name
	location, _ := time.LoadLocation("Asia/Shanghai") // UTC+8 时区
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", s.Start.Format("2006-01-02 15:04:05"), location)
	if err != nil {
		return 0, errors.New("invalid start time format")
	}
	g.Start = startTime
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", s.End.Format("2006-01-02 15:04:05"), location)
	if err != nil {
		return 0, errors.New("invalid end time format")
	}
	if endTime.Before(startTime) {
		return 0, errors.New("end time must be greater than start time")
	}
	g.Start = startTime
	g.End = endTime
	g.End = s.End
	g.Desc = s.Desc
	if id != nil {
		if g.ID != id[0] {
			return 0, errors.New("Game not exists")
		}
		err = db.Save(&g).Error
	} else {
		if g.ID != 0 {
			return 0, errors.New("Game already exists")
		}
		err = db.Create(&g).Error
	}

	if err != nil {
		return 0, err
	}
	return g.ID, nil
}
