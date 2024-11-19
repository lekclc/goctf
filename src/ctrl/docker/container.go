package docker

type Container struct {
	ID    string
	Image string
	Port  map[int]int
}
