package plot

import "fmt"
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
	Position []*Coordinates
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

//func NewPlot(xy *Coordinates) *Plot {
//	g := chart.Chart{
//		Series: []chart.Series{
//			chart.ContinuousSeries{
//				XValues: xy.X,
//				YValues: xy.Y,
//			},
//		},
//	}
//	var p *Plot = new(Plot)
//	p.Position = xy
//	p.Graph = &g
//	return p
//}

func NewPlot(title string) *Plot {
	var g *chart.Chart = new(chart.Chart)
	g.Width, g.Height = chart.DefaultChartWidth, chart.DefaultChartHeight
	var p *Plot = new(Plot)
	p.Graph = g
	return p
}

func (p *Plot) Add(xy *Coordinates, algo string) {
	var cs *chart.ContinuousSeries = new(chart.ContinuousSeries)
	cs.XValues = xy.X
	cs.YValues = xy.Y
	cs.Name = algo
	p.Graph.Series = append(p.Graph.Series, cs)
	fmt.Println(len(p.Graph.Series))
}

func (p *Plot) AddLegend() {
	p.Graph.Elements = []chart.Renderable{
		chart.LegendLeft(p.Graph),
	}
}
