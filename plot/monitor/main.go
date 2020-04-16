package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Monitor struct {
	Mem       []float64
	CPU       []float64
	MaxRecord int
	Lock      sync.Mutex
}

func NewMonitor(max int) *Monitor {
	return &Monitor{
		MaxRecord: max,
	}
}

func (m *Monitor) Collect() {
	mem, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	cpu, err := cpu.Percent(500*time.Millisecond, false)
	if err != nil {
		log.Fatal(err)
	}

	m.Lock.Lock()
	defer m.Lock.Unlock()

	m.Mem = append(m.Mem, mem.UsedPercent)
	m.CPU = append(m.CPU, cpu[0])
}

func (m *Monitor) Run() {
	for {
		m.Collect()
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) WriteTo(w io.Writer) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	cpuData := make(plotter.XYs, len(m.CPU))
	for i, p := range m.CPU {
		cpuData[i].X = float64(i + 1)
		cpuData[i].Y = p
	}

	memData := make(plotter.XYs, len(m.Mem))
	for i, p := range m.Mem {
		memData[i].X = float64(i + 1)
		memData[i].Y = p
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	cpuLine, err := plotter.NewLine(cpuData)
	if err != nil {
		log.Fatal(err)
	}
	cpuLine.Color = plotutil.Color(1)

	memLine, err := plotter.NewLine(memData)
	if err != nil {
		log.Fatal(err)
	}
	memLine.Color = plotutil.Color(2)

	p.Add(cpuLine, memLine)

	p.Legend.Add("cpu", cpuLine)
	p.Legend.Add("mem", memLine)

	p.X.Min = 0
	p.X.Max = float64(m.MaxRecord)
	p.Y.Min = 0
	p.Y.Max = 100

	wc, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		log.Fatal(err)
	}
	wc.WriteTo(w)
}

var monitor = NewMonitor(50)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
}

func image(w http.ResponseWriter, r *http.Request) {
	monitor.WriteTo(w)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/image", image)

	go monitor.Run()

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
