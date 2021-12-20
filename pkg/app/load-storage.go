package app

import (
	"github.com/kenesparta/paackgo/shared/infrastructure/persistence"
)

func (a *App) loadStorage() {
	var err error

	a.fileRepo, err = persistence.NewFileRepository(&a.variables)
	if err != nil {
		a.newLogger.ErrorLogger.Fatalln(err)
	}

	a.memoryRepo, err = persistence.NewMemoryRepository(&a.variables)
	if err != nil {
		a.newLogger.ErrorLogger.Fatalln(err)
	}
}
