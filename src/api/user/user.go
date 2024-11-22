package user

type User struct{}

func NewUser() *User {
	return &User{}
}

type UserInfo struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Passwd string `form:"passwd" json:"passwd" binding:"required"`
	Admin  bool   `form:"admin" json:"admin"`
}
