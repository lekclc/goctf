package docker

import (
	cfg "src/config"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type Container struct {
	ID    string
	Image string
	Port  map[int]int
}

func (c *Container) Create(ID string) {
	cfg.Docker.ContainerCreate(
		cfg.Ctx,
		&container.Config{},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		&v1.Platform{},
		"",
	)
}
