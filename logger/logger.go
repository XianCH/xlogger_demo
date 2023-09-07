package logger

import (
	"log"
	"os"
)

var (
	infoLogger *log.Logger
	logOut     *os.File
)

func SetFile(file string) {
	var err error
	logOut, err = os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	} else {
		infoLogger = log.New(logOut, "[INFO]", log.LstdFlags)

	}
}

func Info(format string, v ...any) {
	infoLogger.Printf(format)
}
