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

type Queue struct {
	first *node
	last  *node
	N     int
}

func (q *Queue) isEmpty() bool {
	return q.N == 0
}

func (q *Queue) size() int {
	return q.N
}

func (q *Queue) enqueue(val interface{}) {
	var last *node = new(node)
	var oldlast *node = new(node)
	oldlast = q.last
	last.item = val
	last.next = nil
	q.last = last
	if q.isEmpty() {
		q.first = last
	} else {
		fmt.Println(oldlast)
		oldlast.next = q.last
		fmt.Println("there")
	}
	q.N++
	fmt.Println("enqueue")
}

func (q *Queue) dequeue() interface{} {
	var oldfirst *node = new(node)
	oldfirst = q.first
	q.first = q.first.next
	if q.isEmpty() {
		q.last = nil
	}
	q.N--
	return oldfirst.item
}

func dynamicCheck() bool {
	var x *Queue = new(Queue)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Queue ok")
	for scanner.Scan() {
		item := scanner.Text()
		if item == "" {
			break
		} else if item == "." {
			fmt.Printf("dequeue: %v <-- length: %v\n", x.dequeue(), x.size())
		} else {
			var n *node = new(node)
			n.item = item
			x.enqueue(n)
		}

	}

	return x.isEmpty() == true
}

func main() {
	dynamicCheck()
}
