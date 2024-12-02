package team_

import (
	con "src/const"
	"src/database"
	"strconv"
)

func (s *Team) Create() (int, string, error) {
	// create team
	db := con.Db.Db
	var team database.Team
	db.Where("name = ?", s.Name).First(&team)
	if team.ID != 0 {
		return 2, "", nil
	}
	team.Name = s.Name
	team.Leader = s.Leader
	team.Desc = s.Desc
	key := "key" //按照队伍名称和时间生成key,类似邀请码的功能
	team.Key = key

	var leader database.User
	leader.Name = s.Leader
	db.Where("name = ?", s.Leader).First(&leader)
	if leader.ID == 0 {
		return 3, "", nil
	}
	team.MemberNum = 1
	db.Create(&team)
	leader.Team += strconv.Itoa(int(team.ID)) + ","
	leader.TeamNum++ //没什么用,但是感觉可能用得到

	err := db.Save(&leader).Error
	if err != nil {
		return 1, "", err
	}

	return 0, key, nil
}
