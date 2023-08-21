package common

type Common struct {
	API      API      `yaml:"api"`
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
	MakefilePath string `yaml:"makefilePath"`
}

type Log struct {
	LogPath string `yaml:"logPath"`
}

type API struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Git struct {
}
