package bumerge

import (
	"fmt"
	"github.com/ntkjer/sedgewick/ds/queue"
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

func (s Sorter) Sort(input []interface{}) {
	var Q *queue.Queue = new(queue.Queue)
	for _, item := range input {
		var qtmp *queue.Queue = new(queue.Queue)
		qtmp.Enqueue(item)
		Q.Enqueue(qtmp)
	}

	for Q.Size() != 1 {
		qFirst, qSecond := Q.Peek(), Q.Peek()
		if qFirst == nil || qSecond == nil {
			break
		}
		var qA *queue.Queue = new(queue.Queue)
		qA.Enqueue(Q.Dequeue())
		var qB *queue.Queue = new(queue.Queue)
		qB.Enqueue(Q.Dequeue())
		qPrime := s.merge(qA, qB)
		Q.Enqueue(qPrime)
	}

}

func exchange(input []interface{}, i, j int) {
	tmp := input[i]
	input[i] = input[j]
	input[j] = tmp
}

func (s Sorter) merge(qA, qB *queue.Queue) *queue.Queue {
	if qA.Size() != qB.Size() {
		panic("mismatched input queue sizes")
	}
	var qOut *queue.Queue = new(queue.Queue)
	itemA, itemB := qA.Peek(), qB.Peek()
	for !qA.IsEmpty() && !qB.IsEmpty() {
		if s.Less(itemA, itemB) {
			qOut.Enqueue(qA.Dequeue())
		} else {
			qOut.Enqueue(qB.Dequeue())
		}
	}

	for !qA.IsEmpty() {
		qOut.Enqueue(qA.Dequeue())
	}
	for !qB.IsEmpty() {
		qOut.Enqueue(qB.Dequeue())
	}

	return qOut
}
