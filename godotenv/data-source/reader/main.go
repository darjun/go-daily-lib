package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	file, _ := os.OpenFile(".env", os.O_RDONLY, 0666)
	myEnv, err := godotenv.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])

	buf := &bytes.Buffer{}
	buf.WriteString("name: awesome web @buffer")
	buf.Write([]byte{'\n'})
	buf.WriteString("version: 0.0.1")
	myEnv, err = godotenv.Parse(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])
}
