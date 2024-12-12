package game_

import (
	"errors"
	con "src/const"
	"src/database"
	"strings"
)

func (s *Game) Join(team uint, name string) error {
	db := con.Db.Db

	var g database.Game
	var t database.Team
	err := db.Where("id = ?", s.ID).First(&g).Error
	if err != nil {
		return err
	}
	err = db.Where("id = ?", team).First(&t).Error
	if err != nil {
		return err
	}
	if t.Leader != name {
		return errors.New("not team leader")
	}
	if t.GameID != 0 {
		return errors.New("team already in game")
	}
	members := strings.Split(t.Member, ",")
	members = members[0 : len(members)-1]
	members = append(members, t.Leader)
	for _, member := range members {
		if IsUserArealdyInGame(member, g.ID) {
			return errors.New("user already in game")
		}
	}

	t.GameID = g.ID
	err = db.Save(&t).Error
	return err
}

func IsUserArealdyInGame(name string, gameid uint) bool {
	db := con.Db.Db
	var u database.User
	err := db.Where("name = ?", name).First(&u).Error
	if err != nil {
		return true
	}
	teamids := strings.Split(u.Team, ",")
	for _, teamid := range teamids {
		if teamid != "" {
			var team database.Team
			err = db.Where("id = ?", teamid).First(&team).Error
			if err != nil {
				return true
			}
			if team.GameID == gameid {
				return true
			}
		}
	}
	return false
}
