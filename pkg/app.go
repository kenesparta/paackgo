package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kenesparta/paackgo/config"
	"github.com/kenesparta/paackgo/shared/infrastructure/api"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	writeTimeout = 5 * time.Second
	readTimeout  = 5 * time.Second
)

type App struct {
	router    *mux.Router
	variables config.VariableConfig
	addr      string
}

func (a *App) loadVariables() {
	switch os.Getenv("GLOBAL_ENV") {
	case "develop":
		a.variables = *config.NewVariableConfig(config.Environment{})
	case "production":
		// The config.AwsSecretsManager{} (not implemented jet)
		a.variables = *config.NewVariableConfig(config.AwsSecretsManager{})
	default:
		log.Fatal("GLOBAL_ENV not set")
	}
	a.addr = fmt.Sprintf(
		"%s:%s",
		a.variables.App.Ip,
		a.variables.App.Port,
	)
}

func (a *App) loadRouter() {
	a.router = mux.NewRouter()
}

func (a *App) loadStorage() {}

func (a *App) loadRoutes() {
	api.InitInMemoryRoutes(a.router)
}

// Initialize Sets the initial configuration for the app
func (a *App) Initialize() {
	a.loadVariables()
	a.loadRouter()
	a.loadStorage()
	a.loadRoutes()
}

// Run Starts the http server
func (a *App) Run() {
	srv := &http.Server{
		Handler:      a.router,
		Addr:         fmt.Sprintf(
			"%s:%s",
			a.variables.App.Ip,
			a.variables.App.Port,
		),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}
	log.Fatal(srv.ListenAndServe())
}
