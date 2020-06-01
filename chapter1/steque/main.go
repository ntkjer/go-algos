package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	item interface{}
	next *node
	prev *node
}

type Stequeue struct {
	first *node
	last  *node
	N     int
}

func (s *Stequeue) push(item interface{}) {
	oldFirst := s.first
	var curr *node = new(node)
	curr.item = item
	curr.next = oldFirst
	s.first = curr

	if oldFirst != nil {
		oldFirst.prev = curr
	} else {
		s.last = curr
	}
	s.N++
}

func (s *Stequeue) isEmpty() bool {
	return s.N == 0
}
func (s *Stequeue) size() int {
	return s.N
}

func (s *Stequeue) pop() interface{} {
	if s.isEmpty() {
		return nil
	}
	result := s.first
	if s.size() > 1 {
		s.first = s.first.next
	} else {
		s.first = nil
	}
	s.first.prev = nil
	s.N--
	return result
}

func (s *Stequeue) enqueue(item interface{}) {
	if s.N == 0 {
		s.push(item)
	} else {
		var newFirst *node = new(node)
		newFirst.item = item
		newFirst.next = s.first
		s.first = newFirst
		s.N++
	}
}

func (s *Stequeue) printLinks() {
	for x := s.first; x != nil; x = x.next {
		fmt.Println(x.item)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s *Stequeue = new(Stequeue)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "pop" {
			fmt.Println(s.pop())
		} else if input == "pl" {
			s.printLinks()
		} else if input == "eq" {
			s.enqueue("enqueue")
		} else {
			s.push(input)
		}
	}
}
