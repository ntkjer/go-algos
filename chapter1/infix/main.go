package main

import (
	"fmt"
)

type Stack struct {
	first *Node
	N     int
}

type Node struct {
	Item string
	next *Node
}

func (s *Stack) isEmpty() bool {
	return s.N == 0

}

func (s *Stack) size() int {
	return s.N

}

func (s *Stack) push(val string) {
	s.N++
	var first *Node = new(Node)
	first.Item = val
	first.next = s.first
	s.first = first
}

func (s *Stack) pop() string {
	result := s.first
	s.first = result.next
	s.N--
	return result.Item

}

func (s *Stack) peek() string {
	return s.first.Item

}

func (s *Stack) isFull() bool {
	return s.size() == s.N

}

func main() {
	//example := "1 + 2 ) * 3 - 4 ) * 5 - 6 ) ) )"
	var input [16]string
	input[0] = "1"
	input[1] = "+"
	input[2] = "2"
	input[3] = ")"
	input[4] = "*"
	input[5] = "3"
	input[6] = "-"
	input[7] = "4"
	input[8] = ")"
	input[9] = "*"
	input[10] = "5"
	input[11] = "-"
	input[12] = "6"
	input[13] = ")"
	input[14] = ")"
	input[15] = ")"
	fmt.Println(getInfix(input))
}

func getInfix(input [16]string) string {
	var operands *Stack = new(Stack)
	var operators *Stack = new(Stack)
	for _, item := range input {
		if item == "(" {
		} else if item == "+" || item == "-" || item == "/" || item == "*" {
			fmt.Println("pushed " + item)
			operators.push(item)
		} else if item == ")" {
			fmt.Println("popping paren " + item)
			operator := operators.pop()
			val2 := operands.pop()
			val1 := operands.pop()
			subExpr := "( " + val1 + " " + operator + val2 + " )"
			operands.push(subExpr)
		} else {
			operands.push(item)
		}
	}
	return operands.pop()
}
