package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/ntkjer/sedgewick/sorts/insertion"
	"github.com/ntkjer/sedgewick/sorts/selection"
	"github.com/ntkjer/sedgewick/sorts/shell"

	"github.com/ntkjer/sedgewick/utils/chart"
	"github.com/ntkjer/sedgewick/utils/stopwatch"
)

func timeTrial(algo string, input []interface{}) float64 {
	var watch *stopwatch.Stopwatch = new(stopwatch.Stopwatch)
	watch.Start()
	if algo == "insertion" {
		var s *insertion.Sorter = new(insertion.Sorter)
		s.Sort(input)
	} else if algo == "selection" {
		var s *selection.Sorter = new(selection.Sorter)
		s.Sort(input)
	} else if algo == "shell" {
		var s *shell.Sorter = new(shell.Sorter)
		s.Sort(input)
	} else {
		panic("not supported yet! please add your sort implementation.")
	}

	watch.Stop()
	return watch.ElapsedTimeSeconds()
}

func randomFloats(min, max float64, n int) []interface{} {
	res := make([]interface{}, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func timedRandomInput(algo string, n, trials int) float64 {
	total := 0.0
	rand.Seed(time.Now().UnixNano())
	input := randomFloats(0.0, 100.01, n)
	for t := 0; t < trials; t++ {
		total += timeTrial(algo, input)
	}
	return total
}

func step(n int) []float64 {
	steps := 10.0
	res := make([]float64, int(steps)+1)
	total := 0.0
	for i := 0; i < len(res); i++ {
		res[i] = total
		total += float64(n) / steps
		fmt.Println(total)
	}
	return res
}

func generatePlotPoints(algoA, algoB string, n, trials int) ([]float64, []float64, []float64) {
	//for a large n, lets step the number into a slice of floats as our x axis
	xaxis := step(n)
	fmt.Println(xaxis)

	aTimes, bTimes := make([]float64, n), make([]float64, n)

	for i, num := range xaxis {
		time1 := timedRandomInput(algoA, int(num), trials)
		time2 := timedRandomInput(algoB, int(num), trials)
		aTimes[i] = time1
		bTimes[i] = time2
	}

	return xaxis, aTimes, bTimes
}

func main() {
	algoA, algoB, n, t := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	N, _ := strconv.Atoi(n)
	trials, _ := strconv.Atoi(t)

	time1 := timedRandomInput(algoA, N, trials)
	time2 := timedRandomInput(algoB, N, trials)
	//fmt.Println(time1, time2)

	ratio := time2 / time1

	fmt.Printf("For %d random floats\n      %s is", N, algoA)
	fmt.Printf("  %.1f times faster than %s.\n", ratio, algoB)

	xaxis, a, b := generatePlotPoints(algoA, algoB, N, trials)
	c1 := plot.NewCoordinates(xaxis, a)
	p := plot.NewPlot("algo runtimes")
	c2 := plot.NewCoordinates(xaxis, b)
	p.Add(c1, algoA)
	p.Add(c2, algoB)
	p.AddLegend()
	p.Output("graph.png")

}
