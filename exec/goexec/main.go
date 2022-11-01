package main

import (
	"fmt"

	"github.com/darjun/goexec"
)

func main() {
	fmt.Println(goexec.CombinedOutputString("cal", nil, goexec.WithEnv("LANG", "en_US.UTF-8")))
}
