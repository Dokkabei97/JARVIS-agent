package model

type Common struct {
	Type     string   `yaml:"type"`
	Script   []Script `yaml:"script,omitempty"`
	Makefile Makefile `yaml:"makefile,omitempty"`
	Log      []Log    `yaml:"log,omitempty"`
}

type Script struct {
	ScriptType string `yaml:"scriptType"`
	ScriptPath string `yaml:"scriptPath"`
}

type Makefile struct {
	MakefilePath   string   `yaml:"makefilePath"`
	MakefileTarget []string `yaml:"makefileTarget"`
}

type Log struct {
	logPath string `yaml:"logPath"`
}
