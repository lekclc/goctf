package challenge_

import (
	con "src/const"
	"src/database"
)

func (s *Challenge) AddFlag() uint {
	// 对于固定flag的题目
	//token, 题目名称, name, flag []string
	//添加flag
	var c database.Challenge
	db := con.Db.Db
	db.Where("name = ?", s.Name).First(&c)
	if c.ID == 0 {
		// 题目不存在
		return 1
	}
	c.Flags += s.Flags + ","
	err := db.Save(&c).Error
	if err != nil {
		// 添加flag失败
		return 2
	}
	return 0
}
