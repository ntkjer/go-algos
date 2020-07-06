package sst

import "reflect"

type SequentialSearchST struct {
	first *Node
}

type Node struct {
	key   interface{}
	value interface{}
	next  *Node
}

func New(key, value interface{}, next *Node) *Node {
	var n *Node
	n.key = key
	n.value = value
	n.next = next
	return n
}

func (s *SequentialSearchST) Get(key interface{}) interface{} {
	for x := s.first; x != nil; x = x.next {
		if equals(key, x.value) {
			return x.value
		}
	}
	return nil
}

func (s *SequentialSearchST) Size() int {
	count := 0
	for curr := s.first; curr != nil; curr = curr.next {
		count++
	}
	return count
}

func (s *SequentialSearchST) isEmpty() bool {
	return s.Size() == 0
}

func (s *SequentialSearchST) Keys() []interface{} {
	result := make([]interface{}, s.Size())
	i := 0
	for x := s.first; x != nil; x = x.next {
		result[i] = x.value
		i++
	}
	return result
}

func (s *SequentialSearchST) Delete(key interface{}) {
	if s.isEmpty() {
		return
	}
	for curr := s.first; curr != nil; curr = curr.next {
		if equals(curr.key, key) {
			curr = nil
			curr = curr.next
		}
	}
}

func (s *SequentialSearchST) Put(key, value interface{}) {
	for x := s.first; x != nil; x = x.next {
		//Search hit, update k : newVal
		if equals(key, x.key) {
			x.value = value
			return
		}
	}
	//Search miss, add to front of ST
	s.first = New(key, value, s.first)
}

func equals(x, y interface{}) bool {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		panic("Error: mismatched inputs")
	}
	switch x.(type) {
	case int:
		a, b := x.(int), y.(int)
		if a == b {
			return true
		}
	case float64:
		a, b := x.(int), y.(int)
		if a == b {
			return true
		}
	case string:
		a, b := x.(string), y.(string)
		if a == b {
			return true
		}
	default:
		panic("unhandled type")
	}
	return false

}
