package docker

import (
	"io"
	cfg "src/config"

	"github.com/docker/docker/api/types/image"
)

type Image struct {
	ID   string
	Name string
	Port map[int]int
}

func NewImage() *Image {
	return &Image{}
}

func (i *Image) Get_info_by_name(name string) {
	a, _, err := cfg.Docker.ImageInspectWithRaw(cfg.Ctx, name)
	if err != nil {
		panic(err)
	}
	i.ID = a.ID
	i.Name = a.RepoTags[0]
}

func (i *Image) Pull_image(name string) {
	res, err := cfg.Docker.ImagePull(cfg.Ctx, name, image.PullOptions{})
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(io.Discard, res)
	if err != nil {
		panic(err)
	}
}

func (i *Image) Build_image() {
	//TODO
}
