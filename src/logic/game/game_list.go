package game_

import (
	"fmt"
	con "src/const"
	"src/database"
	"time"
)

type GameInfo struct {
	Name  string
	Desc  string
	Start time.Time
	End   time.Time
	ID    uint
}

func (g *Game) GameList() (map[string]GameInfo, error) {
	var res = make(map[string]GameInfo)

	var games []database.Game
	db := con.Db.Db

	// 查询所有 game 数据
	if err := db.Find(&games).Error; err != nil {
		return res, err
	}

	// 将查询结果添加到返回值中
	for _, game := range games {
		res[fmt.Sprintf("%d", game.ID)] = GameInfo{
			Name:  game.Name,
			Desc:  game.Desc,
			Start: game.Start,
			End:   game.End,
			ID:    game.ID,
		}
	}

	return res, nil
}
