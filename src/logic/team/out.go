package team_

import (
	"errors"
	con "src/const"
	"src/database"
	"strconv"
	"strings"
)

func (s *Team) Out(name string) error {
	db := con.Db.Db
	var user database.User
	var team database.Team
	err := db.Where("name = ? ", s.Name).First(&team).Error
	if err != nil || team.Leader != s.Leader {
		return err
	}
	err = db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return err
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
		return errors.New("user not found")
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
		return errors.New("user update err")
	}
	db.Save(&team)
	db.Save(&user)
	return nil
}
