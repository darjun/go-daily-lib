package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger.Info().Str("foo", "bar").Msg("hello world")
}