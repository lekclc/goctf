package challenge

type Challenge struct {
	Active    bool
	Name      string
	MaxScore  uint
	ImageID   uint
	DoneNum   uint
	Score     uint
	FileName  string
	ImageName string
	Flags     string
	Hints     string
}

func NewChallenge() *Challenge {
	return &Challenge{}
}
