package unionfind

import (
	"fmt"
	"testing"

	"github.com/ntkjer/sedgewick/chapter1/stopwatch"
)

func TestInitSize(t *testing.T) {
	uf := NewUF(19)
	uf.Union(0, 1)
	uf.Union(0, 2)
	uf.Union(0, 3)
	uf.Union(6, 7)
	uf.Union(8, 9)
	uf.Union(6, 8)
	uf.Union(0, 6)
	uf.Union(10, 11)
	uf.Union(10, 12)
	uf.Union(10, 13)
	uf.Union(10, 14)
	uf.Union(10, 15)
	uf.Union(10, 16)
	uf.Union(10, 17)
	uf.Union(10, 18)
	var w *stopwatch.Stopwatch = new(stopwatch.Stopwatch)
	w.Start()
	uf.Union(0, 10)
	w.Stop()
	expected := 3
	got := uf.Count()
	if expected != got {
		t.Errorf("expected %d, got %d\n", expected, got)
	} else {
		fmt.Println(w.ElapsedTime())
	}
}
