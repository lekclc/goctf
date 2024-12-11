package challenge_

import (
	"errors"
	con "src/const"
	"src/database"
	"strconv"
	"strings"
)

func (s *Challenge) FlagSubmit(gameid uint, challengeid uint, teamid uint, name string) error {
	db := con.Db.Db
	var c database.Challenge
	var t database.Team
	err := db.Where("id = ?", teamid).First(&t).Error
	if err != nil {
		return errors.New("team not found")
	}
	isuser := true
	if t.Leader != name {
		isuser = false
		members := strings.Split(t.Member, ",")
		for _, v := range members {
			if v == name {
				isuser = true
				break
			}
		}
	}
	if !isuser {
		return errors.New("you are not a member of the team")
	}
	err = db.Where("id = ?", challengeid).First(&c).Error
	if err != nil {
		return errors.New("challenge not found")
	}
	var flags []string
	if !c.Active {
		flags = strings.Split(c.Flags, ",")

	} else if c.Active {
		//激活的题目
		var con database.Container
		err = db.Where("challenge_id = ? and team_id = ? and game_id =?", c.ID, teamid, gameid).First(&con).Error
		if err != nil {
			return errors.New("container not found")
		}
		flags = strings.Split(con.Flag, ",")
	}
	for _, v := range flags {
		if v == s.Flags {
			//正确
			t.Challenge = t.Challenge + strconv.Itoa(int(c.ID)) + ","
			t.Score = t.Score + c.Score
			c.DoneNum = c.DoneNum + 1
			if c.Score > 100 {
				c.Score = c.Score - 10
			}
			//更新分数
			db.Save(&t)
			db.Save(&c)
			return nil
		} else {
			return errors.New("flag error")
		}
	}
	return errors.New("error_47")

}
