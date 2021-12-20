package app

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/kenesparta/paackgo/config"
	"github.com/kenesparta/paackgo/logger"
	"github.com/kenesparta/paackgo/shared/infrastructure/persistence"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	router     *mux.Router
	variables  config.VariableConfig
	fileRepo   *persistence.FileRepository
	memoryRepo *persistence.MemoryRepository
	newLogger  *logger.NewLogger
}

// Initialize Sets the initial configuration for the app
func (a *App) Initialize() {
	a.loadVariables()
	a.loadLogger()
	a.loadStorage()
	a.loadRoutes()
}

// Run Starts the http server
func (a *App) Run() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	ctxCancel, cancel := context.WithCancel(context.Background())

	go func() {
		osCall := <-signalChan
		a.newLogger.ErrorLogger.Printf("system call: %v\n", osCall)
		cancel()
	}()

	if err := a.loadServer(ctxCancel); err != nil {
		a.newLogger.ErrorLogger.Printf("failed to serve: %v\n", err)
	}
}
