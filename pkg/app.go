package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kenesparta/paackgo/config"
	"github.com/kenesparta/paackgo/shared/infrastructure/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	writeTimeout   = 2 * time.Second
	readTimeout    = 2 * time.Second
	generalTimeout = 2 * time.Second
	idleTimeout    = 10 * time.Second
)

type App struct {
	router    *mux.Router
	variables config.VariableConfig
}

func (App) loadMainEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (a *App) loadVariables() {
	a.loadMainEnv()
	switch os.Getenv("GLOBAL_ENV") {
	case "develop":
		a.variables = *config.NewVariableConfig(config.Environment{})
	case "production":
		// The config.AwsSecretsManager{} (not implemented jet)
		a.variables = *config.NewVariableConfig(config.AwsSecretsManager{})
	default:
		log.Fatal("GLOBAL_ENV not set")
	}
	a.variables = *config.NewVariableConfig(config.Environment{})
}

func (a *App) loadRoutes() {
	a.router = mux.NewRouter()
	api.InitInMemoryRoutes(a.router)
}

func (a *App) loadStorage() {}

// getServer Creates the main HTTP server
func (a *App) getServer() *http.Server {
	return &http.Server{
		Handler: http.TimeoutHandler(a.router, generalTimeout, ""),
		Addr: fmt.Sprintf(
			":%s",
			a.variables.App.Port,
		),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}
}

// Initialize Sets the initial configuration for the app
func (a *App) Initialize() {
	a.loadVariables()
	a.loadStorage()
	a.loadRoutes()
}

// Run Starts the http server
func (a *App) Run() {
	srv := a.getServer()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-done

	ctxTimeOut, cancel := context.WithTimeout(context.Background(), generalTimeout)
	defer cancel()

	if err := srv.Shutdown(ctxTimeOut); err != nil && err != context.Canceled {
		log.Fatal("shutting down the server gracefully")
	}
}
