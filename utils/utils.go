package utils

import (
	"JARVIS-agent/common"
	"JARVIS-agent/premise"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Common    common.Common     `yaml:"common,omitempty"`
	OnPremise premise.OnPremise `yaml:"onPremise,omitempty"`
	Docker    string            `yaml:"docker,omitempty"`
	Kube      string            `yaml:"kube,omitempty"`
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
