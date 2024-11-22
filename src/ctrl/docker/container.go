package docker

type Container struct {
	ID      string
	Port    map[int]int
	ImageID string
}
