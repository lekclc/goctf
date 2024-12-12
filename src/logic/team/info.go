package team_

import (
	"errors"
	con "src/const"
	"src/database"
)

func (s *Team) Info() (map[any]any, error) {
	// token,team
	// mode, all or one
	var t database.Team

	t.Name = s.Name
	db := con.Db.Db
	db.Where("id = ?", s.ID).First(&t)
	if t.ID == 0 {
		return nil, errors.New("error")
	}
	// 查找所有具有相同 game_id 的团队，并按 score 排序
	var teams []database.Team
	db.Where("game_id = ?", t.GameID).Order("score desc").Find(&teams)

	// 计算当前团队的排名
	rank := 1
	for _, team := range teams {
		if team.ID == t.ID {
			break
		}
		rank++
	}
	return map[any]any{
		"id":        t.ID,
		"name":      t.Name,
		"leader":    t.Leader,
		"members":   t.Member,
		"desc":      t.Desc,
		"score":     t.Score,
		"challenge": t.Challenge,
		"gameid":    t.GameID,
		"key":       t.Key,
		"rank":      rank,
	}, nil
}
