package logger

import (
	"fmt"
	"github.com/kenesparta/paackgo/config"
	"log"
	"os"
	"time"
)

type NewLogger struct {
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
}

func (l *NewLogger) Init(v config.VariableConfig) {
	t := time.Now()
	file, err := os.OpenFile(
		fmt.Sprintf(
			"%s/%d-%d-%d.txt",
			v.LogsDir,
			t.Month(),
			t.Day(),
			t.Year(),
		), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	l.InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
