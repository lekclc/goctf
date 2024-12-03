package team_

import (
	"errors"
	con "src/const"
	"src/database"
)

func (s *Team) Info() (map[any]any, error) {
	// token,team
	// mode, all or one
	var t database.Team
	t.Name = s.Name
	db := con.Db.Db
	db.Where("id = ?", s.ID).First(&t)
	if t.Name == "" {
		return nil, errors.New("error")
	}
	if t.ID == 0 {
		return nil, errors.New("error")
	}
	return map[any]any{
		"name":      t.Name,
		"leader":    t.Leader,
		"members":   t.Member,
		"desc":      t.Desc,
		"score":     t.Score,
		"challenge": t.Challenge,
		"gameid":    t.GameID,
		"key":       t.Key,
	}, nil

	// teamname, teamleader, teammembers[], teamdescription, teamscore, teamrank
}
