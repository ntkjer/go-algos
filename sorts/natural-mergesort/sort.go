package nattymerge

import (
	"fmt"
	"reflect"
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
	s.NaturalSort()
}

func (s Sorter) NaturalSort() {
	if len(s.items) <= 1 {
		return
	}
	offset := 0
	for {
		stop1 := s.findSortedSequence(offset)
		if stop1 == len(s.items)-1 {
			if offset == 0 {
				return
			} else {
				offset = 0
				continue
			}
		}
		stop2 := s.findSortedSequence(stop1 + 1)
		s.merge(s.items, offset, stop1, stop2)
		offset = stop2 + 1

		if stop2 == len(s.items)-1 {
			offset = 0
		}
	}
}

//fancy insertion sort?
func (s Sorter) NaturalSort2() {
	if len(s.items) <= 1 {
		return
	}
	start := 0
	stop1 := s.findSortedSequence(start)
	var stop2 int
	for stop1 < len(s.items)-1 {
		stop2 = s.findSortedSequence(stop1 + 1)

		s.merge(s.items, start, stop1, stop2)
		stop1 = stop2
	}
}

func (s Sorter) findSortedSequence(pos int) int {
	for pos < len(s.items)-1 && s.Less(s.items[pos], s.items[pos+1]) {
		pos = pos + 1
	}
	return pos
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
