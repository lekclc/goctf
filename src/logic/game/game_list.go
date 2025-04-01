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
		fmt.Println(err)
		return res, err
	}
	// 将查询结果添加到返回值中
	for _, game := range games {
		location, _ := time.LoadLocation("Asia/Shanghai") // UTC+8 时区
		startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", game.Start.Format("2006-01-02 15:04:05"), location)
		g.Start = startTime
		endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", game.End.Format("2006-01-02 15:04:05"), location)
		res[fmt.Sprintf("%d", game.ID)] = GameInfo{
			Name:  game.Name,
			Desc:  game.Desc,
			Start: startTime,
			End:   endTime,
			ID:    game.ID,
		}
	}
	fmt.Println("GameList:", res)

	return res, nil
}
