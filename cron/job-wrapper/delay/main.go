package main

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/robfig/cron/v3"
)

type delayJob struct {
	count int32
}

func (d *delayJob) Run() {
	time.Sleep(2 * time.Second)
	atomic.AddInt32(&d.count, 1)
	log.Printf("%d: hello world\n", atomic.LoadInt32(&d.count))
}

func main() {
	c := cron.New()
	c.AddJob("@every 1s", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&delayJob{}))
	c.Start()

	time.Sleep(10 * time.Second)
}
