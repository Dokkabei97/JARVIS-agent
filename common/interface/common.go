package _interface

type CommonService interface {
	GetLog(path string) (string, error)
	GetScript(path string) (string, error)
	ExecuteScript(path string, arguments string) (string, error)
	GetMakefile(path string) (string, error)
	ExecuteMakefile(path string, arguments string) (string, error)
}
