package service

import (
	_interface "JARVIS-agent/docker/interface"
	"JARVIS-agent/utils"
	"context"
	"github.com/docker/docker/api/types"
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
	dockerCli, err := d.utils.GetDockerClient()
	if err != nil {
		return nil, err
	}

	containers, err := dockerCli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	var containerList []string
	for _, container := range containers {
		for _, name := range container.Names {
			containerList = append(containerList, name)
		}
	}

	return containerList, nil
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

func (d dockerService) PullImage(imageName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d dockerService) GetImageList() ([]string, error) {
	//TODO implement me
	panic("implement me")
}
