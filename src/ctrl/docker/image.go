package docker

import (
	"io"

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

func (d *Docker) Get_info_by_name(name string) Image {
	a, _, err := d.cl.ImageInspectWithRaw(d.ctx, name)
	if err != nil {
		panic(err)
	}
	i := Image{}
	i.ID = a.ID
	i.Name = a.RepoTags[0]
	return i
}

func (d *Docker) Pull_image(name string) {
	res, err := d.cl.ImagePull(d.ctx, name, image.PullOptions{})
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(io.Discard, res)
	if err != nil {
		panic(err)
	}
}

func (d *Docker) Build_image() {
}
