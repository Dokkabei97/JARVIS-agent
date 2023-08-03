package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	OnPremise string `yaml:"onPremise,omitempty"`
	Docker    string `yaml:"docker,omitempty"`
	Kube      string `yaml:"kube,omitempty"`
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
