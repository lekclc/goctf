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

func (c *Config) Save() {
	data, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}
	os.WriteFile("config/config.yaml", data, 0644)
}
