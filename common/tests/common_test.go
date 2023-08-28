package tests

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func TestExecuteMakefile2(t *testing.T) {
	testArg := "Arg=wow Arg2=wow2"
	arguments := strings.Split(testArg, " ")

	arg := []string{"-C", "mt"}
	arg = append(arg, arguments...)

	cmd := exec.Command("make", arg...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf("Output: %s", string(output))
}
