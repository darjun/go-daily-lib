package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AddFieldHook struct {
}

func (AddFieldHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level == zerolog.DebugLevel {
		e.Str("name", "dj")
	}
}

func main() {
	hooked := log.Hook(AddFieldHook{})
	hooked.Debug().Msg("")
}
