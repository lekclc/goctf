package team_

import (
	con "src/const"
	"src/database"
	"strconv"
	"strings"
)

func (t *Team) Join(name string) uint {
	var team database.Team
	db := con.Db.Db
	db.Where("name = ?", t.Name).First(&team)
	if team.Key != t.Key {
		return 2
		// key error
	}
	if team.MemberNum >= 4 {
		return 1
		// team full
	}
	var user database.User
	db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		return 3
		// 用户不存在
	}
	members := strings.Split(team.Member, ",")
	for _, member := range members {
		if member == name {
			return 4
			// 用户已经在队伍中
		}
	}

	user.Team += t.Name + ","
	user.TeamNum++
	user.Game += strconv.FormatUint(uint64(t.GameID), 10) + ","

	team.MemberNum++
	team.Member += name + ","
	db.Save(&team)

	return 0
	// success
}
