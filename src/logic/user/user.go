package user_

type User struct {
	Uname  string
	Passwd string
	Admin  bool
}

func GetUser(uname string, passwd string) *User {
	return &User{uname, passwd, false}
}
