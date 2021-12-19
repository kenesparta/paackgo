package app

import (
	"github.com/gorilla/mux"
	"github.com/kenesparta/paackgo/shared/infrastructure/route"
)

func (a *App) loadRoutes() {
	a.router = mux.NewRouter()
	route.InitInMemoryRoutes(a.router)
}
