package cfg

import (
	"context"
	"os"

	"github.com/docker/docker/client"
	"gopkg.in/yaml.v3"
)

var Cfg Config

var Docker *client.Client
var Ctx context.Context

func (c *Config) Init() {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(data, c)
}
