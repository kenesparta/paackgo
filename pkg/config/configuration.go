package config

// iConfiguration created to get environment variables from different fonts
type iConfiguration interface {
	Get(key string) string
}

// VariableConfig The configuration variables from the
type VariableConfig struct {
	App struct {
		Port string
		Ip   string
	}
	File struct {
		Name string
	}
	LogsDir string
}

func NewVariableConfig(c iConfiguration) *VariableConfig {
	builder := variableBuilder{}
	builder.iConfig = c
	return builder.
		appPort().
		fileName().
		logsDir().
		build()
}

type variableActions func(*VariableConfig)

// variableBuilder Sets the variables from other sources
type variableBuilder struct {
	actions []variableActions
	iConfig iConfiguration
}

func (vb *variableBuilder) appPort() *variableBuilder {
	vb.actions = append(vb.actions, func(config *VariableConfig) {
		config.App.Port = vb.iConfig.Get("APP_PORT")
	})
	return vb
}

func (vb *variableBuilder) fileName() *variableBuilder {
	vb.actions = append(vb.actions, func(config *VariableConfig) {
		config.File.Name = vb.iConfig.Get("FILE_NAME")
	})
	return vb
}

func (vb *variableBuilder) logsDir() *variableBuilder {
	vb.actions = append(vb.actions, func(config *VariableConfig) {
		config.LogsDir = vb.iConfig.Get("LOGS_DIR")
	})
	return vb
}

func (vb *variableBuilder) build() *VariableConfig {
	vc := VariableConfig{}
	for _, a := range vb.actions {
		a(&vc)
	}
	return &vc
}
