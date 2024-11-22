package team

type Team struct {
	Name     string
	Leader   string
	Members  []string
	Desc     string
	Key      string
	Chllenge map[string]int
	Score    int
}

func NewTeam() *Team {
	return &Team{}
}
