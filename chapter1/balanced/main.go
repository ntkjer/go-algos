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
	s := fixedStack{entries: items, N: len(items)}
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

func main() {
	fmt.Println(isBalanced())
}

func isBalanced() bool {
	scanner := bufio.NewScanner(os.Stdin)
	var s fixedStack
	var result bool
	symbols := make(map[string]string)
	symbols["}"] = "{"
	symbols["]"] = "["
	symbols[")"] = "("

	for scanner.Scan() {
		item := scanner.Text()
		val, ok := symbols[item]
		fmt.Printf("%v is val\n", val)
		if ok {
			tmp := s.pop()
			if tmp != val {
				fmt.Printf("comparing %v and %v ====== currrent stack looks like: %v\n", tmp, val, &s.entries)
				result = false
			} else {
				result = true
			}
		} else {
			s.push(item)
		}
		if !s.isEmpty() {
			result = false
		}
	}
	return result
}
