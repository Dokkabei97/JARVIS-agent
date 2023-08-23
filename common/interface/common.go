package _interface

type CommonService interface {
	GetLog(path string) (string, error)
	GetFile(path string) (string, error)
	UpdateFile(path string, content string) (string, error)
	ExecuteScript(path string, arguments string) (string, error)
	ExecuteMakefile(path string, arguments string) (string, error)
}
