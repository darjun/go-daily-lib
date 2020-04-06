package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)

	totalPercent, _ := cpu.Percent(3*time.Second, false)
	perPercents, _ := cpu.Percent(3*time.Second, true)
	fmt.Printf("total percent:%v per percents:%v", totalPercent, perPercents)
}
