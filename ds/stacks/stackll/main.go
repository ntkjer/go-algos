package main

import (
	"bufio"
	"fmt"
	"os"
)

type fixedStack struct {
	first *Node
	N     int
}

type Node struct {
	Item interface{}
	next *Node
}

func (s *fixedStack) isEmpty() bool {
	return s.N == 0
}

func (s *fixedStack) size() int {
	return s.N
}

func (s *fixedStack) push(val interface{}) {
	s.N++
	var first *Node = new(Node)
	first.Item = val
	first.next = s.first
	s.first = first
}

func (s *fixedStack) pop() interface{} {
	result := s.first.Item
	s.first = s.first.next
	s.N--
	return result
}

func (s *fixedStack) peek() interface{} {
	return s.first.Item
}

func (s *fixedStack) isFull() bool {
	return s.size() == s.N
}

func CopyStack(toCopy *fixedStack) *fixedStack {
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
	}
}
