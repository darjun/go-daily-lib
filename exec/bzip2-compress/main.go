package main

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func bzipCompress(d []byte) ([]byte, error) {
	var out bytes.Buffer
	cmd := exec.Command("bzip2", "-c", "-9")
	cmd.Stdin = bytes.NewBuffer(d)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}

	return out.Bytes(), nil
}

func main() {
	data := []byte("hello world")
	compressed, _ := bzipCompress(data)
	r := bzip2.NewReader(bytes.NewBuffer(compressed))
	decompressed, _ := ioutil.ReadAll(r)
	fmt.Println(string(decompressed))
}
