package team

type Team struct {
	TeamName    string
	TeamLeader  string
	TeamMembers []string
	TeamDesc    string
	TeamKey     string
	Chllenge    map[string]int
	Score       int
}

func NewTeam() *Team {
	return &Team{}
}
