package docker

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type Image struct {
	Name string
	ID   string
	Port map[int]int
	cl   *client.Client
	ctx  context.Context
}

func NewImage(name string) *Image {
	cl, ctx := Get_cltx()
	return &Image{Name: name, cl: cl, ctx: ctx}
}

func (i *Image) Get_info_by_name() {
	a, _, err := i.cl.ImageInspectWithRaw(i.ctx, i.Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	i.ID = a.ID
	i.Name = a.RepoTags[0]
}

func (i *Image) Pull_image() {
	res, err := i.cl.ImagePull(i.ctx, i.Name, image.PullOptions{})
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(io.Discard, res)
	if err != nil {
		panic(err)
	}
}

func (i *Image) Run_image(containerName string, cmd []string, env []string, ports map[string]string, addressMappings map[string]string, volumeMappings ...map[string]string) string {
	config := &container.Config{
		Image: i.Name,
		Cmd:   cmd,
		Env:   env,
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{},
		Binds:        []string{},
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
	}

	for hostPort, containerPort := range ports {
		hostConfig.PortBindings[nat.Port(containerPort)] = []nat.PortBinding{
			{
				HostPort: hostPort,
			},
		}
	}

	for hostAddress, containerAddress := range addressMappings {
		hostConfig.PortBindings[nat.Port(containerAddress)] = []nat.PortBinding{
			{
				HostIP:   hostAddress,
				HostPort: containerAddress,
			},
		}
	}

	// 如果提供了路径映射参数，则进行路径映射
	if len(volumeMappings) > 0 {
		for hostPath, containerPath := range volumeMappings[0] {
			hostConfig.Binds = append(hostConfig.Binds, fmt.Sprintf("%s:%s", hostPath, containerPath))
		}
	}

	c, err := i.cl.ContainerCreate(i.ctx, config, hostConfig, nil, nil, containerName)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.ID)
	return c.ID
}
