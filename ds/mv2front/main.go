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
}

type LinkedList struct {
	first *node
	N     int
}

func (n node) getItem() interface{} {
	return n.Item
}

func (l *LinkedList) isEmpty() bool {
	return l.first == nil
}

func (l *LinkedList) size() int {
	return l.N
}

func (l *LinkedList) insertFirst(item interface{}) {
	var first *node = new(node)
	oldfirst := l.first
	first.Item = item
	first.next = oldfirst
	l.first = first
	l.N++
}

func (l *LinkedList) append(item interface{}) {
	for x := l.first; x != nil; x = x.next {
		if x.next == nil {
			var last *node = new(node)
			last.next = nil
			x.next = last
			break
		}
	}
}

func (l *LinkedList) MoveToFront(key interface{}) {
	//move most recently accessed items to front of the list
	if !l.isEmpty() {
		fmt.Println("finding key")
		index := l.find(key)
		fmt.Println("key is ", index)
		if index == -1 {
			fmt.Println("not found and inserting ")
			l.insertFirst(key)
		} else {
			fmt.Println("found and attempting to insert at ", index)
			accessedItem := l.removeN(index)
			l.insertFirst(accessedItem)
			fmt.Printf("inserted %v at %d \n", key, index)
		}
	} else {
		l.insertFirst(key)
	}
}

func (l *LinkedList) find(key interface{}) int {
	result, count := -1, 0
	if !l.isEmpty() {
		for x := l.first; x != nil; x = x.next {
			if x.Item == key {
				fmt.Println("x.item == key")
				result = count
				break
			}
			count++
		}
	}
	return result
}

func (l *LinkedList) remove(key interface{}) {
	index := l.find(key)
	if index == -1 {
		return
	}
	l.deleteAfter(index)
}

func (l *LinkedList) removeN(index int) interface{} {
	return l.delete(index)
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

func (l *LinkedList) max(x *node) int {
	max := -1
	for ; x != nil; x = x.next {
		tmp := x.Item.(int)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func (l *LinkedList) recursiveMax(x *node, max int) int {
	if x == nil {
		return max
	}
	tmp := x.Item.(int)
	if tmp > max {
		max = tmp
	}
	return l.recursiveMax(x.next, max)
}

func (l *LinkedList) deleteLast() {
	if !l.isEmpty() {
		if l.size() == 1 {
			l.first = nil
		} else {
			curr := l.first
			for i := 1; i < l.size()-1; i++ {
				curr = curr.next
			}
			curr.next = nil
		}
		l.N--
	}
}

func (l *LinkedList) addAfter(k int, item interface{}) {
	if k > l.size() || (l.isEmpty() && k > 0) {
		panic("index too high")
	} else {
		i := 0
		for curr := l.first; curr != nil; curr = curr.next {
			if i == k {
				var x *node = new(node)
				x.Item = item
				x.next = curr.next
				curr.next = x
				l.N++
				break
			}
			i++
		}
	}
}

//deletes kth element and returns the underlying val of that node
func (l *LinkedList) delete(k int) interface{} {
	var result *node = new(node)
	if k > l.size() {
		result = nil
	} else {
		curr := l.first
		for i := 0; i < k-1; i++ {
			curr = curr.next
		}
		result = curr.next
		curr.next = curr.next.next
		l.N--
	}
	return result.Item
}

func (l *LinkedList) deleteAfter(k int) bool {
	var result bool
	if k > l.size() {
		fmt.Println("k >>>>")
		result = false
	} else {
		curr := l.first
		for i := 0; i < k-2; i++ {
			curr = curr.next
		}
		curr.next = curr.next.next
		l.N--
		result = true
	}
	return result
}

func (l *LinkedList) printLinks() {
	i := 0
	fmt.Println("attemping to print, size: ", l.size())
	for x := l.first; x != nil; x = x.next {
		fmt.Println(x.Item, i)
		i++
	}
}

func (l *LinkedList) deleteFirst() interface{} {
	var item interface{}
	item = l.first.Item
	l.first = l.first.next
	l.N--
	return item
}

func main() {
	mv2frnt()
}

func mv2frnt() {
	var l *LinkedList = new(LinkedList)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		item := scanner.Text()
		if item == "print" {
			l.printLinks()
		} else {
			l.MoveToFront(item)
		}
	}
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
		} else if item == "." {
			fmt.Printf("deleted first: %v <-- length: %v\n", l.deleteFirst(), l.size())
		} else if item == "da" {
			fmt.Println(l.deleteAfter(l.size()/2 - 1))
			l.printLinks()
		} else if item == "find" {
			fmt.Println(l.find("niels"))
		} else if item == "ia" {
			i := 3
			l.addAfter(i, "baboom")
		} else if item == "pl" {
			l.printLinks()
		} else if item == "max" {
			fmt.Println(l.max(l.first))
		} else if item == "rmax" {
			max := 0
			fmt.Println(l.recursiveMax(l.first, max))
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
	i := 0
	for i < numItems {

		fmt.Println("everywhere")
		if s.first != nil {
			fmt.Printf("%v : <-----, items remaining: %d", s.deleteFirst(), s.N)
		}
		i++
	}
	return true
}
