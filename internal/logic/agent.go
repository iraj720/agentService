package logic

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
)

type Agent struct {
	point        []float32
	agentChan    chan []float32
	logger       *log.Logger
	logingPrefix string
	IsBusy       bool
	dist         []float32
	XMove        float32
	YMove        float32
	lock         *sync.Mutex
}

func NewAgent(logingPrefix string, p []float32, f *os.File) Agent {
	log.SetOutput(f)
	logger := log.Default()
	logingPrefix = fmt.Sprintf("Agent Number(%s)", logingPrefix)
	return Agent{point: p, logger: logger, agentChan: make(chan []float32), logingPrefix: logingPrefix, lock: &sync.Mutex{}}

}

func (a *Agent) log(l string) {
	a.logger.Print(a.logingPrefix, " ", l)
}

func (a *Agent) findDistance() float32 {
	xDiff := a.dist[0] - a.point[0]
	yDiff := a.dist[1] - a.point[1]
	return float32(math.Pow(math.Pow(float64(xDiff), 2)+math.Pow(float64(yDiff), 2), 0.5))

}

func (a *Agent) SetMoves() {
	xDiff := a.dist[0] - a.point[0]
	yDiff := a.dist[1] - a.point[1]
	m := yDiff / xDiff
	a.XMove = 1 / (m + 1)
	a.YMove = m * a.XMove

}
func (a *Agent) findDistance2(dist []float32) float32 {
	xDiff := dist[0] - a.point[0]
	yDiff := dist[1] - a.point[1]
	return float32(math.Pow(math.Pow(float64(xDiff), 2)+math.Pow(float64(yDiff), 2), 0.5))
}

func (a *Agent) UpdatePoint() {
	fmt.Println(a.point)
	if a.dist[0] > a.point[0] {
		a.point[0] += a.XMove
	} else if a.dist[0] < a.point[0] {
		a.point[0] -= a.XMove
	}

	if a.dist[1] > a.point[1] {
		a.point[1] += a.YMove
	} else if a.dist[1] < a.point[1] {
		a.point[1] -= a.YMove
	}
	fmt.Println(a.point)
}
