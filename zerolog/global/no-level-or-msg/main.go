package main

import "github.com/rs/zerolog/log"

func main() {
	log.Log().
		Str("foo", "bar").
		Msg("")
}
