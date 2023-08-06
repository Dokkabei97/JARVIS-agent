package service

import (
	_interface "JARVIS-agent/common/interface"
	"JARVIS-agent/utils"
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
