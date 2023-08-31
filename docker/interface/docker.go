package _interface

// DockerService
// Image Build, Image Remove, Container Remove 같은 기능은 제공 X
type DockerService interface {
	GetContainerList() ([]string, error)
	GetContainerLog(containerId string) (string, error)
	StartContainer(containerId string) (string, error)
	StopContainer(containerId string) (string, error)
	RestartContainer(containerId string) (string, error)
}
