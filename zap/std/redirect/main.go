package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.RedirectStdLog(logger)
	log.Print("redirected standard library")
	undo()

	log.Print("restored standard library")
}
