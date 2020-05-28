package main

import (
	"fmt"
	"strings"
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
	var input []string
	input = append(input, "1")
	input = append(input, "+")
	input = append(input, "2")
	input = append(input, ")")
	input = append(input, "*")
	input = append(input, "3")
	input = append(input, "-")
	input = append(input, "4")
	input = append(input, ")")
	input = append(input, "*")
	input = append(input, "5")
	input = append(input, "-")
	input = append(input, "6")
	input = append(input, ")")
	input = append(input, ")")
	input = append(input, ")")
	fmt.Println(infixToPostfix(getInfix(input)))
}

func getInfix(input []string) string {
	var operands *Stack = new(Stack)
	var operators *Stack = new(Stack)
	for _, item := range input {
		if item == "(" {
		} else if item == "+" || item == "-" || item == "/" || item == "*" {
			operators.push(item)
		} else if item == ")" {
			operator := operators.pop()
			val2 := operands.pop()
			val1 := operands.pop()
			subExpr := "(" + val1 + "" + operator + "" + val2 + ")"
			operands.push(subExpr)
		} else {
			operands.push(item)
		}
	}
	return operands.pop()
}

func infixToPostfix(input string) string {
	in := strings.Split(input, "")
	var operands *Stack = new(Stack)
	var operators *Stack = new(Stack)
	for _, item := range in {
		if item == "(" {
		} else if item == "+" || item == "-" || item == "/" || item == "*" {
			operators.push(item)
		} else if item == ")" {
			operator := operators.pop()
			val2 := operands.pop()
			val1 := operands.pop()
			expr := val1 + " " + val2 + " " + operator
			operands.push(expr)
		} else {
			operands.push(item)
		}
	}
	return operands.pop()
}
