package app

import "github.com/kenesparta/paackgo/logger"

func (a *App) loadLogger() {
	a.newLogger = &logger.NewLogger{}
	a.newLogger.Init(a.variables)
}
