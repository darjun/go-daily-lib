package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info().Msg("hello world")
}
