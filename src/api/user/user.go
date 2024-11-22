package user

type User struct{}

func NewUser() *User {
	return &User{}
}

type Info struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" binding:"required"`
	Admin    bool   `form:"admin" json:"admin"`
}
