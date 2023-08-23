package _interface

type CommonService interface {
	GetLog(path string) (string, error)
	GetFile(path string) (string, error)
	ExecuteScript(path string, arguments string) (string, error)
	UpdateScript(path string, content string) (string, error)
	ExecuteMakefile(path string, arguments string) (string, error)
	UpdateMakefile(path string, content string) (string, error)
}
