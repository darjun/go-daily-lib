package main

import (
	"flag"
	"fmt"
)

var (
	// intflag int
	// boolflag bool
	// stringflag string

	intflag    *int
	boolflag   *bool
	stringflag *string
)

func init() {
	// flag.IntVar(&intflag, "intflag", 0, "int flag value")
	// flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	// flag.StringVar(&stringflag, "stringflag", "default", "string flag value")

	intflag = flag.Int("intflag", 0, "int flag value")
	boolflag = flag.Bool("boolflag", false, "bool flag value")
	stringflag = flag.String("stringflag", "default", "string flag value")
}

func main() {
	flag.Parse()

	// fmt.Println("int flag:", intflag)
	// fmt.Println("bool flag:", boolflag)
	// fmt.Println("string flag:", stringflag)

	fmt.Println("int flag:", *intflag)
	fmt.Println("bool flag:", *boolflag)
	fmt.Println("string flag:", *stringflag)

	fmt.Println(flag.Args())
	fmt.Println("Non-Flag Argument Count:", flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("Argument %d: %s\n", i, flag.Arg(i))
	}

	fmt.Println("Flag Count:", flag.NFlag())
}
