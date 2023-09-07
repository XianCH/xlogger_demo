package main

import (
	"github.com/x14n/xlogger_demo/logger"
)

func main() {
	logger.SetFile("log/hello.log")
	logger.Info("HELLO %s", "x14n")
}
