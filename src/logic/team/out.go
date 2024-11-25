package team_

import (
	con "src/const"
	"src/database"
	"strings"
)

func (s *Team) Out(name string) uint {
	db := con.Db.Db
	var user database.User
	var team database.Team
	db.Where("name = ?", s.Name).First(&team)
	if team.ID == 0 {
		return 1
		// message : team not found
	}
	db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		return 2
		// message : user not found
	}
	members := strings.Split(team.Member, ",")
	for i, member := range members {
		if member == name {
			members = append(members[:i], members[i+1:]...)
			team.Member = strings.Join(members, ",")
			db.Save(&team)
			break
		}
	}
	teams := strings.Split(user.Team, ",")
	for i, team := range teams {
		if team == s.Name {
			teams = append(teams[:i], teams[i+1:]...)
			user.Team = strings.Join(teams, ",")
			db.Save(&user)
			break
		}
	}
	// yes or no
	//if no, message
	return 0
}
