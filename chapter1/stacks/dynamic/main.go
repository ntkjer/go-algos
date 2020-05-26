package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

Welp.. this wont work. Golang expects constant to be just that.. compile time constants.
This means we cant use a direct array representation for the stack.
Using a fixed array requires checking an arrays capacity and extending the array to 2*size for the internal store of items. Arrays are created with the [const]interface{} constructor and the go compiler will complain if the const is of variable type :(
*/
type fixedStack struct {
	N       int
	entries [0]interface{}
}

func (s *fixedStack) isEmpty() bool {
	return s.N == 0
}

func (s *fixedStack) size() int {
	return s.N
}

func (s *fixedStack) resize(max int) {
	const cap = max
	var entries [cap]interface{}
	for i := 0; i < len(s.N); i++ {
		entries[i] = s.entries[i]
	}
	s.entries = entries
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s fixedStack
	fmt.Println("stack ok")
	for scanner.Scan() {
		item := scanner.Text()
		if item != "-" {
			s.push(item)
		} else if !s.isEmpty() {
			fmt.Println(s.pop(), " ")
		}
		fmt.Println("(", s.size(), " left on stack)")
		fmt.Println(&s.entries)
	}
}
