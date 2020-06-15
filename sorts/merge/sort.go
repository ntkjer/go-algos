package merge

import (
	"fmt"
	"reflect"
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

//h value is defined as the h number of sorted subsequences for an input slice, with 4 as default.
func (s Sorter) Sort(input []interface{}) {
	aux := make([]interface{}, len(input))
	s.aux = aux
	s.sort(input, 0, len(input)-1)
}

func (s Sorter) sort(input []interface{}, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	s.sort(input, lo, mid)
	s.sort(input, mid+1, hi)
	s.merge(input, lo, mid, hi)
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
