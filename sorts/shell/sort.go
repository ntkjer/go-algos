package insertion

import (
	"fmt"
	"reflect"
)

//TODO: Generics and Comparable

type Sorter struct {
	items  []interface{}
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
	n := len(input)
	h := 1
	for h < n/3 { // 1, 4, 13, 40, 121
		h = 3*h + 1
	}
	for h >= 1 {
		//h sort slice
		for i := h; i < n; i++ {
			for j := i; j >= h && s.Less(input[j], input[j-h]); j -= h {
				exchange(input, j, j-h)
			}
		}
		h = h / 3
	}

	s.sorted = true
}

func exchange(input []interface{}, i, j int) {
	tmp := input[i]
	input[i] = input[j]
	input[j] = tmp
}