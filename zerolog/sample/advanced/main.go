package main

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	sampled := log.Sample(&zerolog.LevelSampler{
		DebugSampler: &zerolog.BurstSampler{
			Burst:       5,
			Period:      time.Second,
			NextSampler: &zerolog.BasicSampler{N: 100},
		},
	})

	sampled.Debug().Msg("hello world")
}
