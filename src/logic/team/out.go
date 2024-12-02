package team_

import (
	con "src/const"
	"src/database"
	"strconv"
	"strings"
)

func (s *Team) Out(name string) uint {
	db := con.Db.Db
	var user database.User
	var team database.Team
	db.Where("name = ? ", s.Name).First(&team)
	if team.ID == 0 || team.Leader != s.Leader {
		return 1
		// message : team not found
	}
	db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		return 2
		// message : user not found
	}
	members := strings.Split(team.Member, ",")
	is := true
	for i, member := range members {
		if member == name {
			members = append(members[:i], members[i+1:]...)
			team.Member = strings.Join(members, ",")
			team.MemberNum--
			is = false

			break
		}

	}
	if is {
		return 3
	}
	is = true
	teams := strings.Split(user.Team, ",")
	for i, team_ := range teams {
		if team_ == strconv.Itoa(int(team.ID)) {
			teams = append(teams[:i], teams[i+1:]...)
			user.Team = strings.Join(teams, ",")
			user.TeamNum--
			is = false
			break
		}
	}
	if is {
		return 4
	}
	if is {
		return 5
	}
	db.Save(&team)
	db.Save(&user)
	return 0
}
