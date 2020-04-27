package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stderr)
	sublogger := logger.With().
		Str("foo", "bar").
		Logger()
	sublogger.Info().Msg("hello world")
}
