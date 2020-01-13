package main

import (
  "fmt"
  "log"
  "os"
    
  "github.com/jessevdk/go-flags"
)

type Option struct {
  Basic GroupBasicOption `description:"basic type" group:"basic"`
  Slice GroupSliceOption `description:"slice of basic type" group:"slice"`
}

type GroupBasicOption struct {
  IntFlag    int     `short:"i" long:"intflag" description:"int flag"`
  BoolFlag   bool    `short:"b" long:"boolflag" description:"bool flag"`
  FloatFlag  float64 `short:"f" long:"floatflag" description:"float flag"`
  StringFlag string  `short:"s" long:"stringflag" description:"string flag"`
}

type GroupSliceOption struct {
  IntSlice		int			`long:"intslice" description:"int slice"`
  BoolSlice		bool		`long:"boolslice" description:"bool slice"`
  FloatSlice	float64	`long:"floatslice" description:"float slice"`
  StringSlice	string	`long:"stringslice" description:"string slice"`
}

func main() {
  var opt Option
  p := flags.NewParser(&opt, flags.Default)
  _, err := p.ParseArgs(os.Args[1:])
  if err != nil {
    log.Fatal("Parse error:", err)
  }
    
  basicGroup := p.Command.Group.Find("basic")
  for _, option := range basicGroup.Options() {
    fmt.Printf("name:%s value:%v\n", option.LongNameWithNamespace(), option.Value())
  }
	
  sliceGroup := p.Command.Group.Find("slice")
  for _, option := range sliceGroup.Options() {
    fmt.Printf("name:%s value:%v\n", option.LongNameWithNamespace(), option.Value())
  }
}