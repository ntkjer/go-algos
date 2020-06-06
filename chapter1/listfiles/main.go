package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func CopyQueue(toCopy *Queue) *Queue {
	var q *Queue = new(Queue)
	for x := toCopy.first; x != nil; x = x.next {
		q.enqueue(x.item)
	}
	return q
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
		oldlast.next = q.last
	}
	q.N++
}

func (q *Queue) listFiles(f string, depth int) {
	fi, err := os.Stat(f)
	if err != nil {
		fmt.Println(err)
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		files, err := ioutil.ReadDir(f)
		if err != nil {
			log.Fatal(err)
		}
		for _, currf := range files {
			q.listFiles(f+"/"+currf.Name(), depth+1)
		}
	case mode.IsRegular():
		q.enqueueFile(f, depth)
	}
}

func (q *Queue) enqueueFile(f string, depth int) {
	var result strings.Builder
	for i := 0; i < depth; i++ {
		result.WriteString("    ")
	}
	result.WriteString(f)
	q.enqueue(result.String())
}

func testFileQueue() {
	var q *Queue = new(Queue)
	path := os.Args[1]
	q.listFiles(path, 0)
	for x := q.first; x != nil; x = x.next {
		fmt.Println(x.item)
	}
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
	testFileQueue()
}
