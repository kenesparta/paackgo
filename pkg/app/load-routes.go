package app

import (
	"github.com/gorilla/mux"
	"github.com/kenesparta/paackgo/shared/infrastructure/api"
)

func (a *App) loadRoutes() {
	a.router = mux.NewRouter()
	api.InitInMemoryRoutes(a.memoryRepo, a.fileRepo, a.router, a.newLogger)
	api.InitFileStorageRoutes(a.fileRepo, a.router, a.newLogger)
}
