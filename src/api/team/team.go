package team

type Team struct {
	Name     string `form:"team" json:"team" binding:"required"`
	Leader   string `form:"name" json:"name"`
	Members  []string
	Desc     string `form:"desc" json:"desc"`
	Key      string
	Chllenge map[string]int
	Score    uint
	GameID   uint `form:"game_id" json:"game_id"`
}

func NewTeam() *Team {
	return &Team{}
}
