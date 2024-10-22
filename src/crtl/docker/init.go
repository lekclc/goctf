package docker

import (
	"context"
	cfg "src/config"

	"github.com/docker/docker/client"
)

type Docker struct {
}

func NewDocker() *Docker {
	return &Docker{}
}

func (d *Docker) Init() {
	cl, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	cfg.Docker = cl
	cfg.Ctx = ctx
}
