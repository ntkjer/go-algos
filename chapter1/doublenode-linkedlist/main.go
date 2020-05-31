package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	Item interface{}
	next *node
	prev *node
}

type LinkedList struct {
	first *node
	last  *node
	N     int
}

func (l *LinkedList) isEmpty() bool {
	return l.first == nil
}

func (l *LinkedList) size() int {
	return l.N
}

func (l *LinkedList) insertFirst(item interface{}) {
	var newFirst *node = new(node)
	newFirst.Item = item
	if l.isEmpty() {
		newFirst.next = nil
		newFirst.prev = nil
		l.first = newFirst
		l.last = newFirst

	} else {
		newFirst.next = l.first
		newFirst.prev = nil
		l.first.prev = newFirst
		l.first = newFirst
	}
	l.N++
	fmt.Println(l.first.Item, l.last.Item, l.size())
	fmt.Println(l.first.Item, l.first.next, l.first.prev)
}

func (l *LinkedList) insertLast(item interface{}) {
	var newLast *node = new(node)
	newLast.Item = item
	if l.isEmpty() {
		newLast.next = nil
		newLast.next = nil
		l.last = newLast
		l.first = newLast
	}
	newLast.next = nil
	newLast.prev = l.last
	l.last.next = newLast
	l.last = newLast
	l.N++
}

func (l *LinkedList) insertBefore(k int, item interface{}) {
	var curr *node = new(node)
	curr.Item = item

	if k > l.size() || l.isEmpty() {
		panic("k > size of *Linkedlist")
	}

	mid := l.size() / 2

	if k > mid {
		i := l.size()
		for x := l.last; i != mid; x = x.prev {
			if i == k {
				curr.next = x
				curr.prev = x.prev
				x.prev.next = curr
				x.prev = curr
				break
			}
			i--
		}
	} else {
		i := 0
		for x := l.first; i != mid; x = x.next {
			if i == k {
				curr.next = x
				curr.prev = x.prev
				x.prev.next = curr
				x.prev = curr
				break
			}
			i++
		}
	}
}

func (l *LinkedList) insertAfter(k int, item interface{}) {
	var curr *node = new(node)
	curr.Item = item
	i := 0
	for x := l.first; x != nil; x = x.next {
		if i == k {
			curr.prev = x
			curr.next = x.next
			x.next.prev = curr
			x.next = curr
			break
		}
		i++
	}
	l.N++
}

func (l *LinkedList) remove(k int) {
	i := 0
	for x := l.first; x != nil; x = x.next {
		if i == k {
			x.prev.next = x.next
			x.next.prev = x.prev
			x = nil
		}
	}
}

func (l *LinkedList) deleteFirst() {
	if l.size() == 0 {
	}
	l.first = l.first.next
	l.first.prev = nil
	l.N--
}

func (l *LinkedList) deleteLast() {
	if l.size() == 0 {
	}
	l.last = l.last.prev
	l.last.next = nil
	l.N--
}

func reverse(x *node) *node {
	var first *node = x
	var reverse *node = nil
	for first != nil {
		second := first.next
		first.next = reverse
		reverse = first
		first = second
	}
	return reverse
}

func rReverse(first *node) *node {
	if first == nil {
		return nil
	}
	if first.next == nil {
		return first
	}
	second := first.next
	rest := rReverse(second)
	return rest
}

func (l *LinkedList) printLinks() {
	i := 0
	for x := l.first; x != nil; x = x.next {
		fmt.Println(x.Item, i)
		i++
	}
}

func main() {
	dynamicCheck()
}

func dynamicCheck() bool {
	var l *LinkedList = new(LinkedList)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("LinkedList ok")
	for scanner.Scan() {
		item := scanner.Text()
		if item == "" {
			break
		} else if item == "dl" {
			fmt.Println("calling delete last, size is: ", l.size())
			l.deleteLast()
			fmt.Println("called delete last, size is now: ", l.size())
			fmt.Println("head is at -> ", l.first)
		} else if item == "df" {
			l.deleteFirst()
			fmt.Printf("deleted first <-- length: %v\n", l.size())
		} else if item == "remove" {
			fmt.Println("removing val, curr size is:", l.size())
			l.remove(l.size()/2 - 1)
			fmt.Println(l.size())
			l.printLinks()
		} else if item == "ia" {
			i := 3
			l.insertAfter(i, "baboom")
		} else if item == "ib" {
			i := 3
			l.insertBefore(i, "boomboom")
		} else if item == "il" {
			l.insertLast("last")
		} else if item == "pl" {
			l.printLinks()
		} else if item == "reverse" {
			l.first = reverse(l.first)
		} else if item == "rreverse" {
			l.first = rReverse(l.first)
		} else {
			num, _ := strconv.Atoi(item)
			l.insertFirst(num)
			//l.insertFirst(item)
		}
	}
	return l.isEmpty() == true
}

func initCheck() bool {
	var first *node = new(node)
	var second *node = new(node)
	var third *node = new(node)
	var fourth *node = new(node)
	var s *LinkedList = new(LinkedList)
	fmt.Println(s.isEmpty())
	first.Item = "to"
	first.next = second
	fmt.Println("here")
	second.Item = "be"
	second.next = third
	third.Item = "or"
	third.next = fourth
	fourth.Item = "that"
	fmt.Println("there")
	s.first = first
	numItems := 4
	s.N = numItems
	return true
}
