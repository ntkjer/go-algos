package main

import (
	"fmt"
)

type stack struct {
	entries []interface{}
	N       int
}

func newStack(items []interface{}) *stack {
	s := stack{entries: items, N: len(items)}
	s.entries = make([]interface{}, s.N)
	return &s
}

func (s *stack) isEmpty() bool {
	return s.N == 0
}

func (s *stack) size() int {
	return s.N
}

func (s *stack) push(item interface{}) {
	s.N++
	s.entries = append(s.entries, item)
	fmt.Printf("pushed : %v\n", item)
}

func (s *stack) pop() interface{} {
	s.N--
	result := s.entries[s.N]
	s.entries = s.entries[:s.N]
	//	if s.isEmpty() {
	//		result = nil
	//	}
	fmt.Printf("popped : %v\n", result)
	return result
}

func main() {
	var check bool
	check = checkIsBalanced()
	if check {
		fmt.Println("balanced!")
	} else {
		fmt.Println("not balanced!")
	}

}

func checkIsBalanced() bool {
	var s *stack = new(stack)
	var items []string
	items = append(items, "[")
	items = append(items, "{")
	items = append(items, "(")
	items = append(items, ")")
	items = append(items, "}")
	items = append(items, "]")
	return s.isBalanced(items)
}

func (s *stack) isBalanced(sequence []string) bool {
	for _, item := range sequence {
		switch item {
		case "{":
			s.push(item)
			break
		case "(":
			s.push(item)
			break
		case "[":
			s.push(item)
			break
		case "]":
			if s.isEmpty() || s.pop() != "[" {
				return false
			}
			break
		case "}":
			if s.isEmpty() || s.pop() != "{" {
				return false
			}
			break
		case ")":
			if s.isEmpty() || s.pop() != "(" {
				return false
			}
			break
		}
	}
	return s.isEmpty()
}

//func isBalanced() bool {
//	scanner := bufio.NewScanner(os.Stdin)
//	var s stack
//	var result bool
//	symbols := make(map[string]string)
//	symbols["}"] = "{"
//	symbols["]"] = "["
//	symbols[")"] = "("
//
//	for scanner.Scan() {
//		item := scanner.Text()
//		val, ok := symbols[item]
//		fmt.Printf("%v is val\n", val)
//		if ok {
//			tmp := s.pop()
//			if tmp != val {
//				fmt.Printf("comparing %v and %v ====== currrent stack looks like: %v\n", tmp, val, &s.entries)
//				result = false
//			} else {
//				result = true
//			}
//		} else {
//			s.push(item)
//		}
//		if !s.isEmpty() {
//			result = false
//		}
//	}
//	return result
//}
