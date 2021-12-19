package app

import (
	"github.com/joho/godotenv"
	"github.com/kenesparta/paackgo/config"
	"log"
	"os"
)

func (a *App) loadVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	switch os.Getenv("GLOBAL_ENV") {
	case "develop":
		a.variables = *config.NewVariableConfig(config.Environment{})
	case "production":
		// The config.AwsSecretsManager{} (not implemented jet)
		a.variables = *config.NewVariableConfig(config.AwsSecretsManager{})
	default:
		log.Fatal("variable GLOBAL_ENV not set")
	}
}
