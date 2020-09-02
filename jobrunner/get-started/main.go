package main

import (
	"fmt"
	"time"

	"github.com/bamzi/jobrunner"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello, ", g.Name)
}

func main() {
	jobrunner.Start()
	jobrunner.Schedule("@every 5s", GreetingJob{Name: "dj"})

	time.Sleep(10 * time.Second)
}
