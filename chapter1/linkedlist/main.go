package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	item interface{}
	next *node
}

type stack struct {
	first *node
	N     int
}

func (s *stack) isEmpty() bool {
	return s.first == nil
}

func (s *stack) size() int {
	return s.N
}

func (s *stack) push(item interface{}) {
	var first *node = new(node)
	oldfirst := s.first
	first.item = item
	first.next = oldfirst
	s.first = first
	s.N++
	fmt.Printf("pushed %v to the stack and is %v next with %v items.\n", s.first, s.first.next, s.size())
}

func (s *stack) pop() interface{} {
	var item interface{}
	item = s.first.item
	s.first = nil
	s.first = s.first.next
	s.N--
	return item
}

func main() {
	dynamicCheck()
}

func dynamicCheck() bool {
	var s *stack = new(stack)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("stack ok")
	for scanner.Scan() {
		item := scanner.Text()
		if item == "" {
			break
		} else if item == "." {
			fmt.Printf("popped: %v <-- length: %v\n", s.pop(), s.size())
		} else {
			var n *node = new(node)
			n.item = item
			s.push(n)
		}

	}

	return s.isEmpty() == true

}

func initCheck() bool {
	var first *node = new(node)
	var second *node = new(node)
	var third *node = new(node)
	var fourth *node = new(node)
	var s *stack = new(stack)
	fmt.Println(s.isEmpty())
	first.item = "to"
	first.next = second
	fmt.Println("here")
	second.item = "be"
	second.next = third
	third.item = "or"
	third.next = fourth
	fourth.item = "that"
	fmt.Println("there")
	s.first = first
	numItems := 4
	s.N = numItems
	i := 0
	for i < numItems {

		fmt.Println("everywhere")
		if s.first != nil {
			fmt.Printf("%v : <-----, items remaining: %d", s.pop(), s.N)
		}
		i++
	}
	return true
}
