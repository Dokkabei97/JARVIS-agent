package service

import (
	_interface "JARVIS-agent/common/interface"
	"JARVIS-agent/utils"
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
	//TODO implement me
	panic("implement me")
}

func (c commonService) GetScript(path string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c commonService) ExecuteScript(path string, arguments string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c commonService) UpdateScript(path string, content string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c commonService) GetMakefile(path string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c commonService) ExecuteMakefile(path string, arguments string) (string, error) {
	argument := strings.Split(arguments, " ")

	cmd := exec.Command("make", argument...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	//TODO implement me
	panic("implement me")
}

func (c commonService) UpdateMakefile(path string, content string) (string, error) {
	//TODO implement me
	panic("implement me")
}
