package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type panicJob struct {
	count int
}

func (p *panicJob) Run() {
	p.count++
	if p.count == 1 {
		panic("oooooooooooooops!!!")
	}

	fmt.Println("hello world")
}

func main() {
	c := cron.New()
	c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&panicJob{}))
	c.Start()

	time.Sleep(5 * time.Second)
}
