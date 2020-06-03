package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type node struct {
	prev *node
	next *node
	item interface{}
}

type Dequeue struct {
	left  *node
	right *node
	N     int
}

func (d *Dequeue) isEmpty() bool {
	return d.N == 0
}

func (d *Dequeue) size() int {
	return d.N
}

func (d *Dequeue) pushLeft(val interface{}) {
	var curr *node = new(node)
	curr.item = val
	if d.isEmpty() {
		d.left = curr
		d.right = curr
	} else {
		curr.next = d.left
		d.left.prev = curr
		d.left = curr
	}
	d.N++
}

func (d *Dequeue) pushRight(val interface{}) {
	var curr *node = new(node)
	curr.item = val
	if d.isEmpty() {
		d.left = curr
		d.right = curr
	} else {
		curr.prev = d.right
		d.right.next = curr
		d.right = curr
	}
	d.N++
}

func (d *Dequeue) popLeft() *node {
	result := d.left
	if result == nil {
		return nil
	}
	if d.left == d.right {
		d.left = nil
		d.right = nil
		d.N--
		return result
	}
	d.left = d.left.next
	d.left.prev = nil
	d.N--
	return result
}

func (d *Dequeue) popRight() *node {
	result := d.right
	if result == nil {
		return nil
	}
	if d.left == d.right {
		d.left = nil
		d.right = nil
		d.N--
		return result
	}
	d.right = d.right.prev
	d.right.next = nil
	d.N--
	return result
}

func (d *Dequeue) printLinks() {
	for x := d.left; x != nil; x = x.next {
		fmt.Println(x.item)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var d *Dequeue = new(Dequeue)

	for scanner.Scan() {
		item := scanner.Text()
		r := rand.Intn(2)
		if item == "popleft" {
			val := d.popLeft()
			fmt.Println(val.item)
		} else if item == "popright" {
			val := d.popRight()
			fmt.Println(val.item)
		} else if item == "print" {
			d.printLinks()
		} else if r == 0 {
			d.pushLeft(item)
		} else if r == 1 {
			d.pushRight(item)
		} else if item == "pl" {
			d.pushLeft("left")
		} else {
			d.pushRight("right")
		}
	}
}
