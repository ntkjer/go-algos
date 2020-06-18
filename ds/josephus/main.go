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
	oldlast := q.last
	last.item = val
	last.next = nil
	q.last = last
	if q.isEmpty() {
		q.first = last
	} else {
		oldlast.next = last
	}
	q.N++
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

func (q *Queue) josephus(n int, m int) {
	for i := 0; i < n; i++ {
		q.enqueue(i)
	}
	for n > 0 {
		for i := 1; i < m; i++ {
			q.enqueue(q.dequeue())
		}
		fmt.Println(q.dequeue())
		n--
	}
}

func isPowerOfTwo(x int) bool {
	return (x != 0) && (x&(x-1)) == 0
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
	var x *Queue = new(Queue)
	x.josephus(7, 2)
}
