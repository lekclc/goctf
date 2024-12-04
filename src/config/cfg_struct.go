package cfg

type host struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type db struct {
	Uname  string `yaml:"uname"`
	Passwd string `yaml:"passwd"`
	Dbname string `yaml:"dbname"`
	Ip     string `yaml:"ip"`
	Port   int    `yaml:"port"`
	Str    string `yaml:"str"`
	Cname  string `yaml:"cname"`
}

type admin struct {
	Uname  string `yaml:"uname"`
	Passwd string `yaml:"passwd"`
}
type str struct {
	Key string `yaml:"key"`
}

type Config struct {
	Host  host  `yaml:"host"`
	Db    db    `yaml:"db"`
	Admin admin `yaml:"admin"`
	Str   str   `yaml:"str"`
}

type UserInfo struct {
	Uname string
	Admin bool
}
