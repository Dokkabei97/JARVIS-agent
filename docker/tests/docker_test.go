package tests

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"testing"
)

// 윈도우 환경에서는 테스트는 WSL에서 실행해야 함

func getDockerClient() *client.Client {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return dockerClient
}

func TestGetDockerClient(t *testing.T) {
	client := getDockerClient()

	if client == nil {
		fmt.Println("Get Docker Client Test => False")
		fmt.Printf("error: %v", "client is nil")
	} else {
		fmt.Println("Get Docker Client Test => True")
	}
}

func TestGetContainerList(t *testing.T) {
	client := getDockerClient()

	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	for _, container := range containers {
		for _, name := range container.Names {
			fmt.Println(name)
		}
	}
}
