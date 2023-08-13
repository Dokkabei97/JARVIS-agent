package _interface

type CommonService interface {
	GetLog(path string) (string, error)
	GetScript(path string) (string, error)
	ExecuteScript(path string) (string, error)
	GetMakefile(path string) (string, error)
	ExecuteMakefile(path string, target string) (string, error)
}
