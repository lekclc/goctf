package game_

import (
	"errors"
	con "src/const"
	"src/database"
)

func (g *Game) Show() (map[string]interface{}, error) {
	db := con.Db.Db
	var game database.Game
	var res = make(map[string]interface{})
	db.Where("id = ?", g.ID).First(&game)
	if game.Name == "" || game.ID == 0 {
		return res, errors.New("Game not found")
	}

	// 查询所有满足条件的 challenge
	var challenges []database.Challenge
	db.Where("game_id = ?", game.ID).Find(&challenges)

	//遍历challenges
	for i := 0; i < len(challenges); i++ {
		challenges[i].Flags = ""
	}
	// 将 challenges 信息添加到结果中
	res["game"] = game
	res["challenges"] = challenges

	return res, nil
}
