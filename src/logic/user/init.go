package user_

type User_ struct {
	Uname  string
	Passwd string
	Admin  bool
}

func NewUser(uname string, passwd string) *User_ {
	return &User_{uname, passwd, false}
}

func GetJwt() {

}
