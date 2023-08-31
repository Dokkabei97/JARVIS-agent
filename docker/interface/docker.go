package _interface

type DockerService interface {
	GetContainerList() ([]string, error)
	GetContainerInfo(containerId string) (string, error)
	GetContainerLog(containerId string) (string, error)
}
