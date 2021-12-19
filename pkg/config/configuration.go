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
}

func NewVariableConfig(c iConfiguration) *VariableConfig {
	builder := variableBuilder{}
	builder.iConfig = c
	return builder.
		appPort().
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

func (vb *variableBuilder) build() *VariableConfig {
	vc := VariableConfig{}
	for _, a := range vb.actions {
		a(&vc)
	}
	return &vc
}
