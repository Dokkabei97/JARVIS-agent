package _interface

type CommonService interface {
	GetLog(path string) (string, error)
}
