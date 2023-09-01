package utils

import (
	"JARVIS-agent/common"
	"JARVIS-agent/docker"
	"fmt"
	"github.com/docker/docker/client"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Common common.Common `yaml:"common,omitempty"`
	Docker docker.Docker `yaml:"docker,omitempty"`
	Kube   string        `yaml:"kube,omitempty"`
}

func (config *Config) GetYamlFile(filepath string) *Config {
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}

	return config
}

func (config *Config) GetDockerClient() (client.Client, error) {
	if config.Common.Type == "docker" {
		dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			fmt.Printf("error: %v", err)
			return client.Client{}, err
		}

		return *dockerClient, nil
	}
	return client.Client{}, nil
}

func (config *Config) GetKubernetesClient() error {
	if config.Common.Type == "kubernetes" {
		//TODO
	}
	return nil
}
