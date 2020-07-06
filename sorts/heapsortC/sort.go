package heapsort

import (
	"fmt"
	//"github.com/ntkjer/sedgewick/sorts/insertion"
	//"github.com/ntkjer/sedgewick/utils/shuffle"

	"github.com/ntkjer/sedgewick/ds/minheap"
	//"github.com/ntkjer/sedgewick/ds/maxheap"
)

type Sorter struct {
	H      *minheap.MinHeap
	sorted bool
	items  []interface{}
}

func NewSorter(input []interface{}) *Sorter {
	var s *Sorter = new(Sorter)
	s.items = input
	s.sorted = false
	return s
}

func (s Sorter) Show(input []interface{}) {
	for _, val := range input {
		fmt.Println(val)
	}
}

func (s Sorter) Sort(input []interface{}) []interface{} {
	result := []interface{}{}

	mh := minheap.New(input)

	for range input {
		result = append(result, mh.DeleteMin())
	}
	return result
}
