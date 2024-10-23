package docker

import (
	"context"

	"github.com/docker/docker/client"
)

type Docker struct {
	ctx context.Context
	cl  *client.Client
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
	d.ctx = ctx
	d.cl = cl
}
