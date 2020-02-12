package main

import (
	"bytes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	buf := &bytes.Buffer{}
	buf.WriteString("name = awesome web")
	buf.WriteByte('\n')
	buf.WriteString("version = 0.0.1")

	env, err := godotenv.Parse(buf)
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Write(env, "./.env")
	if err != nil {
		log.Fatal(err)
	}
}
