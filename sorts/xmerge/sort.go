package bettermerge

import (
	"fmt"
	"time"

	"github.com/ntkjer/sedgewick/sorts/insertion"
)

//TODO: Generics and Comparable

const (
	CUTOFF = 15
)

type Sorter struct {
	items  []interface{}
	aux    []interface{}
	sorted bool
}

func NewSorter(src []interface{}) *Sorter {
	var s *Sorter = new(Sorter)
	s.items = src
	s.sorted = false
	return s
}

func (s Sorter) IsSorted(src []interface{}) bool {
	for i := 1; i < len(src); i++ {
		if s.Less(src[i], src[i-1]) {
			return false
		}
	}
	return true
}

func (s Sorter) Show(src []interface{}) {
	for _, val := range src {
		fmt.Println(val)
	}
}

func (s Sorter) Sort(src []interface{}) {
	aux := make([]interface{}, len(src))
	s.aux = aux
	s.sort(src, aux, 0, len(src)-1)
}

func (s Sorter) sort(src, dst []interface{}, lo, hi int) {
	if hi <= lo+CUTOFF {
		var x *insertion.Sorter = new(insertion.Sorter)
		x.Sort(src[:hi])
		return
	}

	//Insertion sort for small subarrays will give us a ~10-15% boost
	//	if hi <= 15 {
	//
	//		return
	//	}

	mid := lo + (hi-lo)/2
	s.sort(dst, src, lo, mid)
	s.sort(dst, src, mid+1, hi)

	if s.Less(src[mid], src[mid+1]) {
		return
	}

	s.merge(src, dst, lo, mid, hi)
	fmt.Println("merged")
	fmt.Println(src, "\n   ", dst)

	time.Sleep(1 * time.Second)

}

func (s Sorter) merge(src, dst []interface{}, lo, mid, hi int) {

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			dst[k] = src[j]
			j += 1
		} else if j > hi {
			dst[k] = src[i]
			i += 1
		} else if s.Less(src[j], src[i]) {
			dst[k] = src[j]
			j += 1
		} else {
			dst[k] = src[i]
			i += 1
		}
	}
}

func exchange(src []interface{}, i, j int) {
	tmp := src[i]
	src[i] = src[j]
	src[j] = tmp
}

//hacky way for now
func (s Sorter) Less(x, y interface{}) bool {
	switch x.(type) {
	case int:
		a, b := x.(int), y.(int)
		if a > b {
			return false
		}
	case string:
		a, b := x.(string), y.(string)
		if a > b {
			return false
		}
	case float64:
		a, b := x.(float64), y.(float64)
		if a > b {
			return false
		}
	case float32:
		a, b := x.(float32), y.(float32)
		if a > b {
			return false
		}
	case nil:
		panic("nil types cant compare")
	default:
		panic("unhandled types")
	}
	return true
}
