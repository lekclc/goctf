package team_

import (
	"errors"
	con "src/const"
	"src/database"
	"strconv"
	"strings"
)

func (t *Team) Join(name string) error {
	var team database.Team
	db := con.Db.Db
	db.Where("name = ?", t.Name).First(&team)
	if team.Key != t.Key {
		return errors.New("key error")
		// key error
	}
	if team.GameID != 0 {
		return errors.New("team already in game")
		// team already in game
	}
	if team.MemberNum >= 4 {
		return errors.New("team full")
		// team full
	}
	var user database.User
	db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		return errors.New("user not exist")
		// 用户不存在
	}
	members := strings.Split(team.Member, ",")
	for _, member := range members {
		if member == name {
			return errors.New("user already in team")
			// 用户已经在队伍中
		}
	}

	user.Team += strconv.FormatInt(int64(team.ID), 10) + ","
	user.TeamNum++

	team.MemberNum++
	team.Member += name + ","
	db.Save(&team)
	db.Save(&user)

	return nil
	// success
}
