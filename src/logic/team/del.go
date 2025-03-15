package team_

import (
	"errors"
	con "src/const"
	"src/database"
	"strings"
)

func (s *Team) Del(name string, teamid uint) error {
	db := con.Db.Db
	var u database.User
	var t database.Team

	// 开启事务，确保数据一致性
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查询团队信息
	err := tx.First(&t, "id = ?", teamid).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 检查是否为团队领导者
	if t.Leader != name {
		tx.Rollback()
		return errors.New("you are not the leader")
	}

	// 查询领导者用户信息
	err = tx.First(&u, "name = ?", name).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 获取团队成员列表
	members := strings.Split(strings.Trim(t.Member, ","), ",")
	// 包括领导者在内的所有相关用户
	relatedUsers := append(members, t.Leader)
	// 移除空字符串（如果有）
	var validUsers []string
	for _, member := range relatedUsers {
		if member != "" {
			validUsers = append(validUsers, member)
		}
	}

	// 查询所有相关用户的信息
	var users []database.User
	if len(validUsers) > 0 {
		err = tx.Where("name IN ?", validUsers).Find(&users).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 删除团队记录
	err = tx.Delete(&t).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 更新所有相关用户的 team 和 team_num 字段
	for _, user := range users {
		// 从 team 字段中移除被删除的 teamid
		teamList := strings.Split(strings.Trim(user.Team, ","), ",")
		newTeamList := []string{}
		for _, tid := range teamList {
			if tid != "" && tid != string(rune(teamid)) { // 确保不匹配 teamid
				newTeamList = append(newTeamList, tid)
			}
		}

		// 更新 team 和 team_num
		newTeamStr := strings.Join(newTeamList, ",")
		if len(newTeamList) > 0 {
			newTeamStr += ","
		}
		err = tx.Model(&user).Updates(map[string]interface{}{
			"team":     newTeamStr,
			"team_num": len(newTeamList),
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
