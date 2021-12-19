package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	writeTimeout   = 2 * time.Second
	readTimeout    = 2 * time.Second
	generalTimeout = 2 * time.Second
	idleTimeout    = 10 * time.Second
)

// getServer Obtain the server with additional configurations
func (a *App) getServer() *http.Server {
	return &http.Server{
		Handler: http.TimeoutHandler(a.router, generalTimeout, "Timeout!"),
		Addr: fmt.Sprintf(
			":%s",
			a.variables.App.Port,
		),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}
}

// loadServer Creates the main HTTP server
func (a *App) loadServer(ctx context.Context) (err error) {
	log.Printf("server is running at :%s", a.variables.App.Port)
	srv := a.getServer()

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), generalTimeout)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed: %s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
