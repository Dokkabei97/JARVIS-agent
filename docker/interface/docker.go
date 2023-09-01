package _interface

// DockerService
// 1차안에서는 Image Build, Image Remove, Container Remove 같은 기능은 제공 X
type DockerService interface {
	GetContainerList() ([]string, error)
	GetContainerLog(containerId string) (string, error)
	StartContainer(containerId string) (string, error)
	StopContainer(containerId string) (string, error)
	RestartContainer(containerId string) (string, error)
	PullImage(imageName string) (string, error)
	GetImageList() ([]string, error)
}
