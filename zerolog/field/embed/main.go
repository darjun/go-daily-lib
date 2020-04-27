package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().
		Dict("dict", zerolog.Dict().
			Str("bar", "baz").
			Int("n", 1),
		).Msg("hello world")
}
