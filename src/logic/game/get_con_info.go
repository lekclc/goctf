package game_

import (
	"errors"
	cfg "src/config"
	con "src/const"
	"src/database"
	"strings"
)

type info struct {
	ChallengeID uint
	Nc          string
}

func (s *Game) GetConInfo(gameid uint, teamid uint, name string) ([]info, error) {
	db := con.Db.Db
	var t database.Team

	err := db.Where("id = ?", teamid).First(&t).Error
	if err != nil {
		return nil, err
	}
	members := t.Member
	member := strings.Split(members, ",")
	var is_user = false
	if name == t.Leader {
		is_user = true
	} else {
		for _, v := range member {
			if v == name {
				is_user = true
				break
			}
		}
	}
	if !is_user {
		return nil, errors.New("name is not in the team")
	}
	var c []database.Container
	err = db.Where("game_id = ? AND team_id = ?", gameid, teamid).Find(&c).Error
	if err != nil {
		return nil, err
	}
	var res []info
	for _, v := range c {
		var info info
		info.ChallengeID = v.ChallengeID
		info.Nc = cfg.Cfg.Host.Ip + ":" + v.Port
		res = append(res, info)
	}
	return res, nil
}
