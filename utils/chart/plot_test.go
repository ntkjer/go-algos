package plot

import "os"
import "testing"

//Helper for testing Coordinates
func initCoord() *Coordinates {
	x := []float64{
		1.0, 2.0, 3.0, 4.0, 5.0,
	}
	y := []float64{
		1.0, 2.0, 3.0, 4.0, 5.0,
	}
	return NewCoordinates(x, y)
}

func initPlot() *Plot {
	c := initCoord()
	return NewPlot(c)
}

func TestInitCoordinates(t *testing.T) {
	c := initCoord()
	gotX, gotY := c.X, c.Y
	if len(gotX) != len(gotY) {
		t.Errorf("expected length of x and y axis to be equal")
	}
}

func TestInitPlot(t *testing.T) {
	p := initPlot()
	filename := "test.png"
	p.Output(filename)
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		t.Errorf("%s was not rendered by graph.\n", filename)
	}

}
