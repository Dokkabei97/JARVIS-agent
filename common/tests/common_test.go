package tests

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestExecuteMakefile(t *testing.T) {
	cmd := exec.Command("make")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf("Output: %s", string(output))
}
