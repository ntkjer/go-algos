package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
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
	s.N--
	s.first = result.next
	return result.Item

}

func (s *Stack) peek() string {
	return s.first.Item

}

func (s *Stack) isFull() bool {
	return s.size() == s.N

}

func main() {
	fmt.Println(evalPostfix())
}

func evalPostfix() string {
	var operands *Stack = new(Stack)
	var operators *Stack = new(Stack)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		item := scanner.Text()
		if item == "+" || item == "-" || item == "/" || item == "*" {
			fmt.Println("pushing")
			fmt.Println(item)

			time.Sleep(2 * time.Second)
			operators.push(item)
		} else if operands.size() < 2 {
			fmt.Println("size < 2")
			fmt.Println(item)
			operands.push(item)
		} else {
			fmt.Println("popping here")
			operator := operators.pop()
			val2, _ := strconv.Atoi(operands.pop())
			val1, _ := strconv.Atoi(operands.pop())
			var result string
			if operator == "+" {
				result = strconv.Itoa(val1 + val2)
			} else if operator == "-" {
				result = strconv.Itoa(val1 - val2)
			} else if operator == "/" {
				result = strconv.Itoa(val1 / val2)
			} else if operator == "*" {
				result = strconv.Itoa(val1 * val2)
			}
			operands.push(result)
		}
	}
	return operands.pop()
}
