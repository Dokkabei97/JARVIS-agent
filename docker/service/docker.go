package service

import (
	"JARVIS-agent/docker"
	_interface "JARVIS-agent/docker/interface"
	"JARVIS-agent/utils"
	"context"
	"github.com/docker/docker/api/types"
	"io"
)

type dockerService struct {
	utils utils.Config
}

func NewDockerService(util utils.Config) _interface.DockerService {
	return &dockerService{
		utils: util,
	}
}

func (ds dockerService) GetContainerList() ([]docker.Container, error) {
	dockerCli, err := ds.utils.GetDockerClient()
	if err != nil {
		return nil, err
	}
	defer dockerCli.Close()

	containers, err := dockerCli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	containerList := make([]docker.Container, 0)
	for _, container := range containers {
		containerList = append(containerList, docker.Container{
			Id:   container.ID,
			Name: container.Names[0],
			Image: docker.Image{
				Id:   container.ImageID,
				Name: container.Image,
			},
			Status:  container.Status,
			Created: container.Created,
		})
	}

	return containerList, nil
}

func (ds dockerService) GetContainerLog(container docker.Container, follow bool, tail string) (io.ReadCloser, error) {
	dockerCli, err := ds.utils.GetDockerClient()
	if err != nil {
		return nil, err
	}
	defer dockerCli.Close()

	if container.Id == "" || container.Name == "" {
		return nil, err
	}

	ctx := context.Background()
	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     follow,
		Tail:       tail,
		Timestamps: true,
	}

	containerIdORName := container.Id
	if containerIdORName == "" {
		containerIdORName = container.Name
	}

	logs, err := dockerCli.ContainerLogs(ctx, containerIdORName, options)
	if err != nil {
		return nil, err
	}
	defer logs.Close()

	return logs, nil
}

func (ds dockerService) RunContainer(imageName string, containerName string, port string, options ...string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (ds dockerService) StartContainer(containerId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (ds dockerService) StopContainer(containerId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (ds dockerService) RestartContainer(containerId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (ds dockerService) RemoveContainer(containerId string, isForce bool) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (ds dockerService) PullImage(imageName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (ds dockerService) PushImage(imageName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (ds dockerService) RemoveImage(imageName string, isForce bool) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (ds dockerService) GetImageList() ([]docker.Image, error) {
	dockerCli, err := ds.utils.GetDockerClient()
	if err != nil {
		return nil, err
	}
	defer dockerCli.Close()

	images, err := dockerCli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return nil, err
	}

	imageList := make([]docker.Image, 0)
	for _, image := range images {
		imageList = append(imageList, docker.Image{
			Id:      image.ID,
			Name:    image.RepoTags[0],
			Size:    image.Size,
			Created: image.Created,
		})
	}

	return imageList, nil
}

func (ds dockerService) BuildImage(imageName string, dockerFilePath string) (string, error) {
	//TODO implement me
	panic("implement me")
}
