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
	symbols["{"] = "}"
	symbols["["] = "]"
	symbols["("] = ")"

	for scanner.Scan() {
		item := scanner.Text()
		if item == "" {
			break
		}
		val, ok := symbols[item]
		if ok {
			if s.isEmpty() {
				result = false
			} else {
				tmp := s.pop()
				if tmp == val {
					result = true
				}
			}
		} else {
			s.push(item)
		}
		if s.isEmpty() {
			return result
		}
	}
	return result
}
