package quicksort

import (
	"fmt"
	"reflect"
	//"github.com/ntkjer/sedgewick/sorts/insertion"

	"github.com/ntkjer/sedgewick/utils/shuffle"
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
	fisheryates.Shuffle(input)
	s.sort(input, 0, len(input)-1)
}

func (s Sorter) sort(input []interface{}, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(input, lo, hi)
	s.sort(input, lo, j-1)
	s.sort(input, j+1, hi)
}

// problem 2.3.17
func partitionS(input []interface{}, lo, hi int) int {
	i := lo
	j := hi
	v := input[lo]
	for {
		for Less(input[i], v) {
			if i == hi {
				break
			}
			i = i + 1
		}
		for Less(v, input[j]) {
			if j == lo {
				break
			}
			j = j - 1
		}
		if i >= j {
			break
		}
		exchange(input, i, j)
	}
	exchange(input, lo, j)
	return j
}

func partition(input []interface{}, lo, hi int) int {
	i := lo
	j := hi
	v := input[lo]
	for {
		for Less(input[i], v) {
			if i == hi {
				break
			}
			i = i + 1
		}
		for Less(v, input[j]) {
			if j == lo {
				break
			}
			j = j - 1
		}
		if i >= j {
			break
		}
		exchange(input, i, j)
	}
	exchange(input, lo, j)
	return j
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

func Less(x, y interface{}) bool {
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
