package docker

type Container struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Image   Image  `json:"image"`
	Status  string `json:"status"`
	Created int64  `json:"created"`
}

type Image struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	Created int64  `json:"created"`
}

type Docker struct {
	Containers []Container `json:"containers"`
	Images     []Image     `json:"images"`
}
