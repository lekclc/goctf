package game_

import (
	con "src/const"
	"src/database"
	"time"
)

type Info struct {
	Name  string
	Desc  string
	Start time.Time
	End   time.Time
}

func (s *Game) GameList() (map[any]any, error) {
	res := make(map[any]any)

	var games []database.Game
	db := con.Db.Db

	// 查询所有 game 数据
	if err := db.Find(&games).Error; err != nil {
		return res, err
	}

	// 将查询结果添加到返回值中
	for _, game := range games {
		res[game.ID] = Info{
			Name:  game.Name,
			Desc:  game.Desc,
			Start: game.Start,
			End:   game.End,
		}
	}

	return res, nil
}
