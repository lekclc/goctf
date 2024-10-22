package cfg

type host struct {
	Url  string `yaml:"url"`
	Port int    `yaml:"port"`
}

type Config struct {
	Host host `yaml:"host"`
}
