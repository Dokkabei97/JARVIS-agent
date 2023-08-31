package utils

import (
	"JARVIS-agent/common"
	"JARVIS-agent/docker"
	"fmt"
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

	err = config.validateYaml()
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}

	return config
}

func (config *Config) validateYaml() error {
	if config.Common.Type == "docker" {
		//TODO
	} else if config.Common.Type == "kube" {
		//TODO
	}

	return nil
}

func (config *Config) GetLogPaths() []string {
	var logPaths []string

	for _, log := range config.Common.Log {
		logPaths = append(logPaths, log.LogPath)
	}

	return logPaths
}
