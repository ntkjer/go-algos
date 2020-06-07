package linkedlist

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

func (l *LinkedList) IsEmpty() bool {
	return l.first == nil
}

func (l *LinkedList) Size() int {
	return l.N
}

func (l *LinkedList) InsertFirst(item interface{}) {
	var first *node = new(node)
	oldfirst := l.first
	first.Item = item
	first.next = oldfirst
	l.first = first
	l.N++
}

func (l *LinkedList) Append(item interface{}) {
	for x := l.first; x != nil; x = x.next {
		if x.next == nil {
			var last *node = new(node)
			last.Item = item
			last.next = nil
			x.next = last
			break
		}
	}
}

func (l *LinkedList) Find(key interface{}) int {
	result, count := -1, 0
	if !l.IsEmpty() {
		for x := l.first; x != nil; x = x.next {
			if x.Item == key {
				result = count
				break
			}
			count++
		}
	}
	return result
}

func (l *LinkedList) Remove(key interface{}) {
	index := l.Find(key)
	if index == -1 {
		return
	}
	l.DeleteAfter(index)

}

func Reverse(x *node) *node {
	var first *node = x
	var Reverse *node = nil
	for first != nil {
		second := first.next
		first.next = Reverse
		Reverse = first
		first = second
	}
	return Reverse
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

func (l *LinkedList) RecursiveMax(x *node, max int) int {
	if x == nil {
		return max
	}
	tmp := x.Item.(int)
	if tmp > max {
		max = tmp
	}
	return l.RecursiveMax(x.next, max)
}

func (l *LinkedList) DeleteLast() {
	if !l.IsEmpty() {
		if l.Size() == 1 {
			l.first = nil
		} else {
			curr := l.first
			for i := 1; i < l.Size()-1; i++ {
				curr = curr.next
			}
			curr.next = nil
		}
		l.N--
	}
}

func (l *LinkedList) AddAfter(k int, item interface{}) {
	if k > l.Size() || (l.IsEmpty() && k > 0) {
		panic("index too high")
	} else {
		i := 0
		for curr := l.first; curr != nil; curr = curr.next {
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

func (l *LinkedList) DeleteAfter(k int) bool {
	var result bool
	if k > l.Size() {
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

func (l *LinkedList) PrintLinks() {
	i := 0
	for x := l.first; x != nil; x = x.next {
		fmt.Println(x.Item, i)
		i++
	}
}

func (l *LinkedList) DeleteFirst() interface{} {
	var item interface{}
	item = l.first.Item
	l.first = l.first.next
	l.N--
	return item
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
			fmt.Println("calling delete last, Size is: ", l.Size())
			l.DeleteLast()
			fmt.Println("called delete last, Size is now: ", l.Size())
			fmt.Println("head is at -> ", l.first)
		} else if item == "." {
			fmt.Printf("deleted first: %v <-- length: %v\n", l.DeleteFirst(), l.Size())
		} else if item == "da" {
			fmt.Println(l.DeleteAfter(l.Size()/2 - 1))
			l.PrintLinks()
		} else if item == "Find" {
			fmt.Println(l.Find("niels"))
		} else if item == "ia" {
			i := 3
			l.AddAfter(i, "baboom")
		} else if item == "pl" {
			l.PrintLinks()
		} else if item == "max" {
			fmt.Println(l.max(l.first))
		} else if item == "rmax" {
			max := 0
			fmt.Println(l.RecursiveMax(l.first, max))
		} else if item == "Reverse" {
			l.first = Reverse(l.first)
		} else if item == "rReverse" {
			l.first = rReverse(l.first)
		} else {
			num, _ := strconv.Atoi(item)
			l.InsertFirst(num)
			//l.InsertFirst(item)
		}
	}
	return l.IsEmpty() == true
}

func initCheck() bool {
	var first *node = new(node)
	var second *node = new(node)
	var third *node = new(node)
	var fourth *node = new(node)
	var s *LinkedList = new(LinkedList)
	fmt.Println(s.IsEmpty())
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
			fmt.Printf("%v : <-----, items remaining: %d", s.DeleteFirst(), s.N)
		}
		i++
	}
	return true
}
