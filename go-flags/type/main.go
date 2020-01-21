package main

import (
  "fmt"

  "github.com/jessevdk/go-flags"
)

type Option struct {
  IntFlag         int             `short:"i" long:"int" description:"int flag value"`
  IntSlice        []int           `long:"intslice" description:"int slice flag value"`
  BoolFlag        bool            `long:"bool" description:"bool flag value"`
  BoolSlice       []bool          `long:"boolslice" description:"bool slice flag value"`
  FloatFlag       float64         `long:"float" description:"float64 flag value"`
  FloatSlice      []float64       `long:"floatslice" description:"float64 slice flag value"`
  StringFlag      string          `short:"s" long:"string" description:"string flag value"`
  StringSlice     []string        `long:"strslice" description:"string slice flag value"`
  PtrStringSlice  []*string       `long:"pstrslice" description:"slice of pointer of string flag value"`
  Call            func(string)    `long:"call" description:"callback"`
  IntMap          map[string]int  `long:"intmap" description:"A map from string to int"`
}

func main() {
  var opt Option
  opt.Call = func (value string) {
    fmt.Println("in callback: ", value)
  }
  
  _, err := flags.Parse(&opt)
  if err != nil {
    fmt.Println("Parse error:", err)
    return
  }
  
  fmt.Printf("int flag: %v\n", opt.IntFlag)
  fmt.Printf("int slice flag: %v\n", opt.IntSlice)
  fmt.Printf("bool flag: %v\n", opt.BoolFlag)
  fmt.Printf("bool slice flag: %v\n", opt.BoolSlice)
  fmt.Printf("float flag: %v\n", opt.FloatFlag)
  fmt.Printf("float slice flag: %v\n", opt.FloatSlice)
  fmt.Printf("string flag: %v\n", opt.StringFlag)
  fmt.Printf("string slice flag: %v\n", opt.StringSlice)
  fmt.Println("slice of pointer of string flag: ")
  for i := 0; i < len(opt.PtrStringSlice); i++ {
    fmt.Printf("\t%d: %v\n", i, *opt.PtrStringSlice[i])
  }
  fmt.Printf("int map: %v\n", opt.IntMap)
}