package tests

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestExecuteMakefile(t *testing.T) {
	cmd := exec.Command("bash", "-c", "make")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf("Output: %s", string(output))
}

func TestGetMakefile(t *testing.T) {
	file, _ := os.ReadFile("./Makefile")

	fmt.Printf("%s", string(file))
}

func TestUpdateMakefile(t *testing.T) {

	content := "all:" +
		"\n\t@echo \"Hello World\"" +
		"\nhi:" +
		"\n\t@echo \"Hi World\""

	file, _ := os.Create("./Makefile")
	defer file.Close()

	file.Write([]byte(content))
}
