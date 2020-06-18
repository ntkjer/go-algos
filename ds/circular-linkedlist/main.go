package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type node struct {
	Item interface{}
	next *node
}

type Queue struct {
	last *node
	N    int
}

func (n node) getItem() interface{} {
	return n.Item
}

func (l *Queue) isEmpty() bool {
	return l.last == nil
}

func (l *Queue) size() int {
	return l.N
}

func (q *Queue) enqueue(item interface{}) {
	var curr *node = new(node)
	if q.N == 0 {
		curr.next = curr
	} else if q.N > 0 {
		curr.next = q.last.next
		q.last.next = curr
	}
	curr.Item = item
	q.last = curr
	q.N++
}

func (l *Queue) append(item interface{}) {
	for x := l.last; x != nil; x = x.next {
		if x.next == nil {
			var last *node = new(node)
			last.Item = item
			last.next = nil
			x.next = last
			break
		}
	}
}

func (l *Queue) find(key interface{}) int {
	result, count := -1, 0
	if !l.isEmpty() {
		for x := l.last; x != nil; x = x.next {
			fmt.Println(reflect.TypeOf(key), reflect.TypeOf(x.Item))
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

func (l *Queue) remove(key interface{}) {
	index := l.find(key)
	if index == -1 {
		return
	}
	l.deleteAfter(index)

}

func (l *Queue) max(x *node) int {
	max := -1
	for ; x != nil; x = x.next {
		tmp := x.Item.(int)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func (l *Queue) recursiveMax(x *node, max int) int {
	if x == nil {
		return max
	}
	tmp := x.Item.(int)
	if tmp > max {
		max = tmp
	}
	return l.recursiveMax(x.next, max)
}

func (l *Queue) deleteLast() {
	if !l.isEmpty() {
		if l.size() == 1 {
			l.last = nil
		} else {
			curr := l.last
			for i := 1; i < l.size()-1; i++ {
				curr = curr.next
			}
			curr.next = nil
		}
		l.N--
	}
}

func (l *Queue) addAfter(k int, item interface{}) {
	if k > l.size() || (l.isEmpty() && k > 0) {
		panic("index too high")
	} else {
		i := 0
		for curr := l.last; curr != nil; curr = curr.next {
			if i == k {
				var x *node = new(node)
				x.Item = item
				x.next = curr.next
				fmt.Println(curr.next.Item, x.next.Item, x.Item)
				curr.next = x
				fmt.Println("========================")
				fmt.Println("success!")
				l.N++
				break
			}
			i++

		}
	}
}

func (l *Queue) deleteAfter(k int) bool {
	var result bool
	if k > l.size() {
		fmt.Println("k >>>>")
		result = false
	} else {
		curr := l.last
		for i := 0; i < k-2; i++ {
			curr = curr.next
		}
		curr.next = curr.next.next
		l.N--
		result = true
	}
	return result
}

func (l *Queue) printLinks() {
	i := 0
	x := l.last
	for i != l.N {
		fmt.Println(x.Item, i)
		x = x.next
		i++
	}
}

func (l *Queue) deletelast() interface{} {
	var item interface{}
	item = l.last.Item
	l.last = l.last.next
	l.N--
	return item
}

func main() {
	dynamicCheck()
}

func dynamicCheck() bool {
	var l *Queue = new(Queue)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Queue ok")
	for scanner.Scan() {
		item := scanner.Text()
		if item == "" {
			break
		} else if item == "dl" {
			fmt.Println("calling delete last, size is: ", l.size())
			l.deleteLast()
			fmt.Println("called delete last, size is now: ", l.size())
			fmt.Println("head is at -> ", l.last)
		} else if item == "." {
			fmt.Printf("deleted last: %v <-- length: %v\n", l.deletelast(), l.size())
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
			fmt.Println(l.max(l.last))
		} else if item == "rmax" {
			max := 0
			fmt.Println(l.recursiveMax(l.last, max))
		} else {
			num, _ := strconv.Atoi(item)
			l.enqueue(num)
			//l.enqueue(item)
		}
	}
	return l.isEmpty() == true
}

func initCheck() bool {
	var last *node = new(node)
	var second *node = new(node)
	var third *node = new(node)
	var fourth *node = new(node)
	var s *Queue = new(Queue)
	fmt.Println(s.isEmpty())
	last.Item = "to"
	last.next = second
	fmt.Println("here")
	second.Item = "be"
	second.next = third
	third.Item = "or"
	third.next = fourth
	fourth.Item = "that"
	fmt.Println("there")
	s.last = last
	numItems := 4
	s.N = numItems
	i := 0
	for i < numItems {

		fmt.Println("everywhere")
		if s.last != nil {
			fmt.Printf("%v : <-----, items remaining: %d", s.deletelast(), s.N)
		}
		i++
	}
	return true
}
