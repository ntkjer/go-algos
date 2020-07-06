package quicksort3way

import (
	"fmt"
	"reflect"

	"github.com/ntkjer/sedgewick/utils/shuffle"
	//"github.com/ntkjer/sedgewick/sorts/insertion" //for cutoff point
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
	lt, i, gt := lo, lo+1, hi
	v := input[lo]
	for i <= gt {
		cmp := Compare(input[i], v)
		if cmp < 0 {
			exchange(input, lt, i)
			lt += 1
			i += 1
		} else if cmp > 0 {
			exchange(input, i, gt)
			gt -= 1
		} else {
			i += 1
		}
	}
	s.sort(input, lo, lt-1)
	s.sort(input, gt+1, hi)
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

func Compare(x, y interface{}) int {
	var result int
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		panic("mismatched input types for args x and y")
	}
	switch x.(type) {
	case int:
		a, b := x.(int), y.(int)
		if a > b {
			result = 1
		} else if a < b {
			result = -1
		} else {
			result = 0
		}
	case string:
		a, b := x.(string), y.(string)
		if a > b {
			result = 1
		} else if a < b {
			result = -1
		} else {
			result = 0
		}
	case float64:
		a, b := x.(float64), y.(float64)
		if a > b {
			result = 1
		} else if a < b {
			result = -1
		} else {
			result = 0
		}
	case float32:
		a, b := x.(float32), y.(float32)
		if a > b {
			result = 1
		} else if a < b {
			result = -1
		} else {
			result = 0
		}
	default:
		panic("unhandled type")
	}

	return result
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
