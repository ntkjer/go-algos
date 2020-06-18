package main

import (
	"bufio"
	"fmt"
	"os"
)

type fixedStack struct {
	entries []interface{}
	N       int
}

func newStack(items []interface{}) *fixedStack {
	s := fixedStack{entries: items, N: len(items) - 1}
	s.entries = make([]interface{}, s.N)
	return &s
}

func (s *fixedStack) isEmpty() bool {
	return s.N == 0
}

func (s *fixedStack) size() int {
	return s.N
}

func (s *fixedStack) push(item interface{}) {
	s.N++
	s.entries = append(s.entries, item)
}

func (s *fixedStack) pop() interface{} {
	s.N--
	result := s.entries[s.N]
	s.entries = s.entries[:s.N]
	return result
}

func (s *fixedStack) peek() interface{} {
	return s.entries[s.N-1]

}

//Technically this will always be true since N== len(entries)
//Also this doesnt really make sense in the context of slices as the internal collection of items, since they grow automatically.
func (s *fixedStack) isFull() bool {
	return s.size() == s.N
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s fixedStack
	fmt.Println("stack ok")
	for scanner.Scan() {
		item := scanner.Text()
		if item != "-" {
			s.push(item)
			fmt.Println(s.peek(), "its working")
		} else if !s.isEmpty() {
			fmt.Println(s.pop(), " ")
		}
		fmt.Println("(", s.size(), " left on stack)")
		fmt.Println(&s.entries)
	}
}
