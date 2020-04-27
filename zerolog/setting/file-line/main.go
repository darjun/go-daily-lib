package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Caller().Logger()
	logger.Info().Msg("hello world")
}
