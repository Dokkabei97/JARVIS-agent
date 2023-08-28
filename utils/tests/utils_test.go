package tests

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

type TestYaml struct {
	Deploy TestDeploy `yaml:"deploy"`
}

type TestDeploy struct {
	Type         string   `yaml:"type"`
	Host         string   `yaml:"host"`
	ScriptPath   string   `yaml:"scriptPath,omitempty"`
	LogPath      string   `yaml:"logPath,omitempty"`
	ComposePath  string   `yaml:"composePath,omitempty"`
	MakePath     string   `yaml:"makePath,omitempty"`
	kubeYamlPath string   `yaml:"kubeYamlPath,omitempty"`
	ScriptPaths  []string `yaml:"scriptPaths,omitempty"`
}

func TestGetYamlFile(t *testing.T) {
	file, _ := os.ReadFile("./test.yml")

	var testYaml TestYaml
	_ = yaml.Unmarshal(file, &testYaml)

	fmt.Printf("type: %s, "+
		"host: %s, "+
		"scriptPath: %s, "+
		"logPath: %s",
		testYaml.Deploy.Type,
		testYaml.Deploy.Host,
		testYaml.Deploy.ScriptPath,
		testYaml.Deploy.LogPath,
	)
}

func TestGetMultiYamlFile(t *testing.T) {
	file, _ := os.ReadFile("./test.yml")

	decoder := yaml.NewDecoder(bytes.NewReader(file))

	var deploy []TestYaml

	for {
		var d TestYaml
		err := decoder.Decode(&d)
		if err != nil {
			break
		}
		deploy = append(deploy, d)
	}

	for _, d := range deploy {
		fmt.Printf("type: %s, "+
			"host: %s, "+
			"scriptPath: %s, "+
			"logPath: %s\n",
			d.Deploy.Type,
			d.Deploy.Host,
			d.Deploy.ScriptPath,
			d.Deploy.LogPath,
		)
	}
}
