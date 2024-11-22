package challenge

type Challenge struct {
	Name      string
	MaxScore  int
	DoneNum   int
	Score     int
	FileName  string
	ImageName string
	Flags     []string
	Hints     []string
}

func NewChallenge() *Challenge {
	return &Challenge{}
}
