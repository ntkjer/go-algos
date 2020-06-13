package plot

import "testing"

func TestInitCoordinates(t *testing.T) {
	x := []float64{
		1.0, 2.0, 3.0, 4.0, 5.0,
	}
	y := []float64{
		1.0, 2.0, 3.0, 4.0, 5.0,
	}
	c := NewCoordinates(x, y)
	gotX, gotY := c.X, c.Y
	if len(gotX) != len(gotY) {
		t.Errorf("expected length of x and y axis to be equal")
	}
}
