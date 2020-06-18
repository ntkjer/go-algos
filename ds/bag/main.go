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

type Bag struct {
	first *node
	N     int
}

func (b *Bag) isEmpty() bool {
	return b.N == 0
}

func (b *Bag) size() int {
	return b.N
}

func (b *Bag) add(val interface{}) {
	var first *node = new(node)
	oldfirst := b.first
	first.item = val
	first.next = oldfirst
}

func dynamicCheck() bool {
	var x *Bag = new(Bag)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Bag ok")
	for scanner.Scan() {
		item := scanner.Text()
		if item == "" {
			break
		} else {
			var n *node = new(node)
			n.item = item
			x.add(n)
		}

	}

	return x.isEmpty() == true
}

func main() {
	dynamicCheck()
}
