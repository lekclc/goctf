package team_

import (
	con "src/const"
	"src/database"
)

func (s *Team) Out(name string) uint {
	db := con.Db.Db
	var user database.User
	var team database.Team
	db.Where("name = ?", s.Name).First(&team)
	db.Where("name = ?", name).First(&user)
	// yes or no
	//if no, message
	return 0
}
