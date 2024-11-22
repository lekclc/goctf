package docker

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
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

func (i *Image) Run_image() {

}
