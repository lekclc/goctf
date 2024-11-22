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

func Get_cltx() (*client.Client, context.Context) {
	cl_, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	ctx_ := context.Background()
	return cl_, ctx_
}

func (d *Docker) Init() {
	d.cl, d.ctx = Get_cltx()
}
