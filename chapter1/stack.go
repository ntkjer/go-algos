package stack

type Stack struct {
	entries []interface{}
	N       int
}

func newStack(items []interface{}) *Stack {
	s := Stack{entries: items, N: len(items) - 1}
	s.entries = make([]interface{}, s.N)
	return &s
}

func (s *Stack) IsEmpty() bool {
	return s.N == 0
}

func (s *Stack) Size() int {
	return s.N
}

func (s *Stack) Push(item interface{}) {
	s.N++
	s.entries = append(s.entries, item)
}

func (s *Stack) Pop() interface{} {
	s.N--
	result := s.entries[s.N]
	s.entries = s.entries[:s.N]
	return result
}

func (s *Stack) Peek() interface{} {
	return s.entries[s.N-1]

}

//Technically this will always be true since N== len(entries)
//Also this doesnt really make sense in the context of slices as the internal collection of items, since they grow automatically.
func (s *Stack) isFull() bool {
	return s.Size() == s.N
}
