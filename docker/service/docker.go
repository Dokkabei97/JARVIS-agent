package service

import (
	_interface "JARVIS-agent/docker/interface"
	"JARVIS-agent/utils"
)

type dockerService struct {
	utils utils.Config
}

func NewDockerService(util utils.Config) _interface.DockerService {
	return &dockerService{
		utils: util,
	}
}

func (d dockerService) GetContainerList() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (d dockerService) GetContainerLog(containerId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d dockerService) StartContainer(containerId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d dockerService) StopContainer(containerId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d dockerService) RestartContainer(containerId string) (string, error) {
	//TODO implement me
	panic("implement me")
}
