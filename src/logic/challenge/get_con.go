package challenge_

import (
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"os/exec"
	con "src/const"
	"src/ctrl/utils"
	"src/database"
	"strconv"
	"strings"
	"time"
)

func (s *Challenge) GetCon(chanllengid uint, teamid uint, gammeid uint) (string, error) {

	db := con.Db.Db

	var game database.Game
	var team database.Team
	var challenge database.Challenge
	var Container database.Container
	var image database.Image
	db.Where("challenge_id = ? and team_id = ? and game_id = ?", chanllengid, teamid, gammeid).First(&Container)
	if Container.ID != 0 {
		return Container.Port, nil
	}

	err := db.Where("id = ?", gammeid).First(&game).Error
	if err != nil {
		return "", err
	}
	err = db.Where("id = ?", teamid).First(&team).Error
	if err != nil {
		return "", err
	}
	if game.ID != team.GameID {
		return "", errors.New("gameid not match")
	}
	err = db.Where("id = ?", chanllengid).First(&challenge).Error
	if err != nil {
		return "", err
	}
	if challenge.GameID != game.ID {
		return "", errors.New("gameid not match")
	}
	err = db.Where("challenge_id = ?", challenge.ID).First(&image).Error
	if err != nil {
		return "", err
	}
	Container.ChallengeID = challenge.ID
	Container.TeamID = team.ID
	Container.GameID = game.ID

	ports := image.Port
	expose_int, err := utils.GetUsePort()
	if err != nil {
		return "", err
	}
	expose := strconv.Itoa(expose_int)
	cname := team.Name + strconv.Itoa(time.Now().Nanosecond())
	cmd := exec.Command("docker", "run", "-d",
		"-p", expose+":"+ports,
		"--name", cname,
		image.Name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	res := string(output)[:len(string(output))-1]
	Container.Port = expose
	Container.ContainerID = res
	flag := res + team.Name + challenge.Name + time.Now().String()
	hash := sha256.Sum256([]byte(flag))
	flag_ := fmt.Sprintf("%x", hash)
	flag_ = fmt.Sprintf("%x", md5.Sum([]byte(flag_)))

	Container.Flag = strings.Replace(challenge.Flags, "[[FLAG]]", flag_, -1)
	fmt.Println(Container.Flag)
	db.Save(&Container)

	return expose, nil

}
