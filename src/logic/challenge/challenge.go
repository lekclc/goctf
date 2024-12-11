package challenge_

type Challenge struct {
	Class     string `form:"class" json:"class"`
	Active    bool   `form:"active" json:"active"`
	Name      string `form:"challenge_name" json:"challenge_name" binding:"required"`
	MaxScore  uint   `form:"max_score" json:"max_score"`
	ImageID   uint
	DoneNum   uint
	Score     uint
	FileName  string `form:"file_name" json:"file_name"`
	ImageName string `form:"image_name" json:"image_name"`
	Flags     string `form:"flags" json:"flags"`
	Desc      string `form:"desc" json:"desc"`
	GameID    uint
	Port      string
	ID        uint
}

func GetChallenge() *Challenge {
	return &Challenge{}
}
