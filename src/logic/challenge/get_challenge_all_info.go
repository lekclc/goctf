package challenge_

import (
	con "src/const"
	"src/database"
)

func (s *Challenge) GetChallengeAllInfo(challengeid uint) (map[any]any, error) {
	db := con.Db.Db
	var c database.Challenge

	err := db.Where("id = ?", challengeid).First(&c).Error
	if err != nil {
		return nil, err
	}
	res := make(map[any]any)
	res["flags"] = c.Flags
	res["image"] = c.ImageName

	return res, nil

}
