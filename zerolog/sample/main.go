package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	sampled := log.Sample(&zerolog.BasicSampler{N: 10})

	for i := 0; i < 20; i++ {
		sampled.Info().Msg("will be logged every 10 message")
	}
}
