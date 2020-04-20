package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/roylee0704/gron"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello ", g.Name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	g1 := GreetingJob{Name: "dj"}
	g2 := GreetingJob{Name: "dajun"}

	c := gron.New()
	c.Add(gron.Every(5*time.Second), g1)
	c.Add(gron.Every(10*time.Second), g2)
	c.Start()

	wg.Wait()
}
