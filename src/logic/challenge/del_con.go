package challenge_

import (
	"errors"
	"os/exec"
	con "src/const"
	"src/database"
	"strings"
)

func (s *Challenge) DelCon(chanllengid uint, teamid uint, gammeid uint, name string) error {
	db := con.Db.Db
	var t database.Team
	err := db.Where("id = ?", teamid).First(&t).Error
	if err != nil {
		return err
	}
	isuser := true
	if name != t.Leader {
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
		return errors.New("name is not in the team")
	}
	var con database.Container
	err = db.Where("challenge_id = ? AND team_id = ? AND game_id = ?", chanllengid, teamid, gammeid).First(&con).Error
	if err != nil {
		return err
	}
	cmd := exec.Command("docker", "stop", con.ContainerID)

	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("docker", "rm", con.ContainerID)
	err = cmd.Run()
	if err != nil {
		return err
	}

	err = db.Unscoped().Delete(&con).Error
	if err != nil {
		return err
	}

	return nil
}
