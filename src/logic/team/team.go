package team_

type Team struct {
	Name     string
	Leader   string
	Members  []string
	Desc     string
	Key      string
	Chllenge map[string]int
	Score    uint
	GameID   uint
}

func GetTeam() *Team {
	return &Team{}
}
