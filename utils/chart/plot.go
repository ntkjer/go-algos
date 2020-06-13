package plot

import "github.com/wcharczuk/go-chart"
import "os"

type Coordinates struct {
	X []float64
	Y []float64
}

func NewCoordinates(x, y []float64) *Coordinates {
	var c *Coordinates = new(Coordinates)
	c.X = x
	c.Y = y
	return c
}

type Plot struct {
	Position *Coordinates
	Graph    *chart.Chart
}

func (p *Plot) Output(in string) {
	f, err := os.Create(in)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	p.Graph.Render(chart.PNG, f)
}

func (p *Plot) Init() {
	p.Graph.Chart.Series.chart.ContinuousSeries.XValues = p.Graph.X
	p.Graph.Chart.Series.chart.ContinuousSeries.YValues = p.Graph.Y
}

func NewPlot(xy *Coordinates) *Plot {
	g := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xy.X,
				YValues: xy.Y,
			},
		},
	}
	var p *Plot = new(Plot)
	p.Position = xy
	p.Graph = g
	return &p
}
