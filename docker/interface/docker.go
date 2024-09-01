package _interface

import (
	"JARVIS-agent/docker"
	"io"
)

type DockerService interface {
	GetContainerList() ([]docker.Container, error)
	GetContainerLog(container docker.Container, follow bool, tail string) (io.ReadCloser, error)
	RunContainer(imageName string, containerName string, port string, options ...string) (string, error)
	StartContainer(containerId string) (string, error)
	StopContainer(containerId string) (string, error)
	RestartContainer(containerId string) (string, error)
	RemoveContainer(containerId string, isForce bool) (string, error)
	PullImage(imageName string) (string, error)
	PushImage(imageName string) (string, error)
	RemoveImage(imageName string, isForce bool) (string, error)
	GetImageList() ([]docker.Image, error)
	BuildImage(imageName string, dockerFilePath string) (string, error)
}
