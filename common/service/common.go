package service

import (
	_interface "JARVIS-agent/common/interface"
	"JARVIS-agent/utils"
	"bufio"
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io"
	"os"
	"os/exec"
	"strings"
)

type commonService struct {
	utils utils.Config
}

func NewCommonService(util utils.Config) _interface.CommonService {
	return &commonService{
		utils: util,
	}
}

func (c commonService) GetLog(path string) (string, error) {
	return realTimeLogMonitor(path)
}

func (c commonService) GetFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return "", err
	}

	return string(file), nil
}

func (c commonService) UpdateFile(path string, content string) (string, error) {
	file, err := os.Create(path)

	if err != nil {
		fmt.Printf("Error creating file: %v", err)
		return "", err
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		fmt.Printf("Error writing file: %v", err)
		return "", err
	}

	return "file updated", nil
}

func (c commonService) ExecuteScript(path string, arguments string) (string, error) {
	argument := getArgs(arguments)

	cmd := exec.Command(path, argument...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func (c commonService) ExecuteMakefile(arguments string) (string, error) {
	path := c.utils.Common.Makefile.MakefilePath
	argument := getArgs(arguments)

	makeCmd := []string{"-C", path}
	makeCmd = append(makeCmd, argument...)

	cmd := exec.Command("make", makeCmd...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func getArgs(arguments string) []string {
	return strings.Split(arguments, " ")
}

func fileWatch(path string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	defer watcher.Close()

	err = watcher.Add(path)
	if err != nil {
		return nil, err
	}

	return watcher, nil
}

func convertEncoding(input io.Reader, output io.Writer, transformer transform.Transformer) error {
	reader := transform.NewReader(input, transformer)
	_, err := io.Copy(output, reader)
	return err
}

func realTimeLogMonitor(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return "", err
	}

	watcher, err := fileWatch(file.Name())

	reader := bufio.NewReader(file)
	for {
		if err != nil {
			if err == io.EOF {
				select {
				case event := <-watcher.Events:
					if event.Op&fsnotify.Write == fsnotify.Write {
						continue
					}
				case err := <-watcher.Errors:
					fmt.Printf("error: %v", err)
				}
			} else {
				fmt.Printf("error: %v", err)
			}
		}

		buf := new(bytes.Buffer)
		err = convertEncoding(reader, buf, korean.EUCKR.NewDecoder())
		if err != nil {
			fmt.Printf("error: %v", err)
		}

		return buf.String(), nil
	}
}
