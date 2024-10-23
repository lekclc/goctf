package cfg

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg Config

func (c *Config) Init() {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(data, c)
}
