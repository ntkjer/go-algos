package heapsort

import (
	"fmt"
	//"github.com/ntkjer/sedgewick/sorts/insertion"
	//"github.com/ntkjer/sedgewick/utils/shuffle"
)

type Sorter struct {
	Q      *MaxPQ
	sorted bool
}

func NewSorter(input []interface{}) *Sorter {
	var s *Sorter = new(Sorter)
	pq := NewMaxPQ(len(input) + 1)
	for i, num := range input {
		pq.pq[i+1] = num
	}
	s.Q = pq
	s.Q.n = len(input) + 1
	s.sorted = false
	return s
}

func (s Sorter) IsSorted(input []interface{}) bool {
	fmt.Println("checking if sorted.. ")
	s.Show(input)
	for i := 2; i < len(input); i++ {
		if Less(input, i+1, i) {
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
	n := len(input)
	for k := n / 2; k >= 1; k-- {
		Sink(input, k, n)
	}
	for n > 1 {
		Exchange(input, 1, n)
		n = n - 1
		Sink(input, 1, n)
	}
}
