package config

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type NewVariableConfigTestSuiteEnv struct {
	suite.Suite
	variables VariableConfig
}

func (suite *NewVariableConfigTestSuiteEnv) SetupTest() {
	_ = os.Setenv("APP_PORT", "8080")
	suite.variables = *NewVariableConfig(Environment{})
}

func (suite *NewVariableConfigTestSuiteEnv) TestNewVariableConfigNotEmptyAppIPort() {
	suite.NotEqual("", suite.variables.App.Port)
}

func TestNewVariableConfigEnv(t *testing.T) {
	suite.Run(t, new(NewVariableConfigTestSuiteEnv))
}
