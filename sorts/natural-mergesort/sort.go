package nattymerge

import (
	"fmt"
	"reflect"
	"time"
	//"github.com/ntkjer/sedgewick/sorts/insertion"
)

//TODO: Generics and Comparable

type Sorter struct {
	items  []interface{}
	aux    []interface{}
	sorted bool
}

func NewSorter(input []interface{}) *Sorter {
	var s *Sorter = new(Sorter)
	s.items = input
	s.sorted = false
	return s
}

func (s Sorter) IsSorted(input []interface{}) bool {
	for i := 1; i < len(input); i++ {
		if s.Less(input[i], input[i-1]) {
			return false
		}
	}
	return true
}

func (s Sorter) Show(input []interface{}) {
	for _, val := range input {
		fmt.Println(val)
	}
}

func (s Sorter) Sort(input []interface{}) {
	aux := make([]interface{}, len(input))
	s.aux = aux
	s.naturalSort()
}

func (s Sorter) naturalSorty(input []interface{}) {
	lo, hi := 0, 0

	mid := s.findSortedSequence(input, lo)
	for {
		fmt.Println(input)
		time.Sleep(1 * time.Second)
		//Find first sorted subsequence from current i
		// Find second subsequence from

		//merge them
		lo = hi
		mid = s.findSortedSequence(input, lo)
		if lo == 0 && mid == len(input)-1 {
			return
		}
		hi = s.findSortedSequence(input, mid)
		if mid == hi && hi == len(input)-1 {
			lo = 0
			hi = 0
			continue
		}
		s.merge(input, lo, mid, hi)

		fmt.Println(lo, mid, hi)
	}
}

func (s Sorter) naturalSort() {
	if len(s.items) <= 1 {
		return
	}

	offset := 0

	for {
		stop1 := s.findSortedSequence(s.items, offset)

		if stop1 == len(s.items)-1 {

			if offset == 0 {
				// Sorting done.
				return
			} else {
				// Current run is an orphan.
				// Possibly handle it in the
				// next merge pass.
				offset = 0
				continue
			}
		}

		stop2 := s.findSortedSequence(s.items, stop1+1)

		s.merge(s.items, offset, stop1, stop2)
		offset = stop2 + 1

		if stop2 == len(s.items)-1 {
			// Proceed to the next merge pass.
			offset = 0
		}
	}
	s.sorted = true
}

func (s Sorter) findSortedSequence(input []interface{}, pos int) int {
	for i := pos + 1; i < len(input); i++ {
		if s.Less(input[i-1], input[i]) {
			return i - 1
		}
	}
	return len(input) - 1
}

func (s Sorter) merge(input []interface{}, lo, mid, hi int) {

	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		s.aux[k] = input[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			input[k] = s.aux[j]
			j += 1
		} else if j > hi {
			input[k] = s.aux[i]
			i += 1
		} else if s.Less(s.aux[j], s.aux[i]) {
			input[k] = s.aux[j]
			j += 1
		} else {
			input[k] = s.aux[i]
			i += 1
		}
	}
}

func exchange(input []interface{}, i, j int) {
	tmp := input[i]
	input[i] = input[j]
	input[j] = tmp
}

//hacky way for now
func (s Sorter) Less(x, y interface{}) bool {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		panic("mismatched input types for args x and y")
	}
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
	default:
		panic("unhandled type")
	}

	return true
}
