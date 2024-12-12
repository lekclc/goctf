package game_

import (
	con "src/const"
	"src/database"
	"strings"
)

type rankinfo struct {
	TeamName  string
	Rank      uint
	Score     uint
	Challenge map[string]bool
}

func (s *Game) GetAllTeamRank() ([]rankinfo, error) {
	db := con.Db.Db
	var g database.Game
	err := db.Where("id = ?", s.ID).First(&g).Error
	if err != nil {
		return nil, err
	}
	var ts []database.Team
	err = db.Where("game_id = ?", s.ID).Order("score DESC").Find(&ts).Error
	if err != nil {
		return nil, err
	}
	var res []rankinfo
	for i, team := range ts {
		cmap := make(map[string]bool)
		challenge := strings.Split(team.Challenge, ",")
		challenge = challenge[0 : len(challenge)-1]
		for _, c := range challenge {
			cmap[c] = true
		}
		res = append(res, rankinfo{
			TeamName:  team.Name,
			Rank:      uint(i + 1),
			Score:     team.Score,
			Challenge: cmap,
		})
	}
	return res, nil
}
