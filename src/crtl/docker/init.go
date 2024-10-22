package docker

import (
	cfg "src/config"

	"github.com/docker/docker/client"
)

func (*Docker) Init() {
	cl, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	cfg.Docker = cl
}
