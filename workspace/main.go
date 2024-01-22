package main

import (
	"github.com/AdiLambe/TestGoLangCodes/workspace/app"
	"github.com/AdiLambe/TestGoLangCodes/workspace/logger"
)

func main() {

	//log.Println("Starting our application...")
	logger.Info("Starting our application...")

	app.Start()
}
