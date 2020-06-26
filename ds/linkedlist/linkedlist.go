package linkedlist

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
	var reverse *node = nil
	for first != nil {
		second := first.next
		first.next = reverse
		reverse = first
		first = second
	}
	return reverse
}

//Reverses a linked list and saves a bit more memory in the process.
func ReversePointers(first **node) *node {
	var result *node = nil
	var current *node = *first

	for current != nil {
		MoveNode(&result, &current)
	}
	return result
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

//Assumest the linkedlist is in sorted order
func (l *LinkedList) RemoveDuplicates(head *node) {
	if head == nil {
		return
	}
	for curr := head; curr != nil; curr = curr.next {
		if curr.next != nil {
			if !Less(curr.next.Item, curr.Item) {
				curr.next = curr.next.next
			}
		}
	}
}

func (l *LinkedList) IsSorted(head *node) bool {
	for curr := head; curr != nil; curr = curr.next {
		next := curr.next
		if next == nil {
			break
		}
		if Less(next.Item, curr) {
			return false
		}
	}
	return true
}

func (l *LinkedList) FrontBackSplit() (*LinkedList, *LinkedList) {
	slow := l.first
	fast := slow
	var back *LinkedList = new(LinkedList)
	var front *LinkedList = new(LinkedList)

	for i := 0; i < l.Size(); i += 1 {
		slow = slow.next
		front.Append(slow)
		bound := i + 3
		for j := i; j < bound; j++ {
			fast = fast.next
			if j == bound-1 {
				back.Append(fast.Item)
			}
		}
	}
	return front, back
}

func (l *LinkedList) Max(x *node) int {
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
				l.N++
				break
			}
			i++

		}
	}
}

// Takes the head of second list and adds it to onto the front of first list
func MoveNode(sourceRef, destRef **node) {
	var newNode *node = *sourceRef
	if newNode == nil {
		return
	}
	*sourceRef = newNode.next //Advance source pointer
	newNode.next = *destRef
	*destRef = newNode //point the ref to newNode instead
}

func AlternatingSplit(source *node, aRef, bRef **node) {
	var aDummy node
	var aTail *node = &aDummy
	var bDummy node
	var bTail *node = &bDummy
	var current *node = source

	aDummy.next = nil
	bDummy.next = nil

	for current != nil {
		MoveNode(&(aTail.next), &current)
		aTail = aTail.next
		if current != nil {
			MoveNode(&(bTail.next), &current)
			bTail = bTail.next
		}
	}
	*aRef = aDummy.next
	*bRef = bDummy.next
}

func AlternatingSplit3(source *node, aRef, bRef **node) {
	var a *node
	var b *node

	var current *node = source
	for current != nil {
		MoveNode(&a, &current)
		if current != nil {
			MoveNode(&b, &current)
		}
	}
	*aRef = a
	*bRef = b
}

//Takes a linked list and returns two sublists of alternating items
func (l *LinkedList) AlternatingSplit2() (*LinkedList, *LinkedList) {
	var firstHalf *LinkedList = new(LinkedList)
	var secondHalf *LinkedList = new(LinkedList)

	first := false
	curr := l.first
	for i := 0; i < l.Size(); i++ {
		if first {
			firstHalf.Append(curr.Item)
		} else {
			secondHalf.Append(curr.Item)
		}
		curr = curr.next
		first = !first
	}
	return firstHalf, secondHalf
}

//Uses dummy nodes but also MoveNode for compactness
func ShuffleMerge(a, b *node) *node {
	var dummy node
	var tail *node = &dummy
	dummy.next = nil

	for {
		if a == nil {
			tail.next = b
			break
		} else if b == nil {
			tail.next = a
			break
		} else {
			MoveNode(&(tail.next), &a)
			tail = tail.next
			MoveNode(&(tail.next), &b)
			tail = tail.next
		}
	}
	return dummy.next
}

func ShuffleMergeLonger(a, b *node) *node {
	var dummy node
	var tail *node = &dummy
	dummy.next = nil

	for {
		if a == nil {
			tail.next = b
			break
		} else if b == nil {
			tail.next = a
			break
		} else {
			tail.next = a
			tail = a
			a = a.next
			tail.next = b
			tail = b
			b = b.next
		}
	}
	return dummy.next
}

//Uses stack space proportional to size of list
func ShuffleMergeRecursive(a, b *node) *node {
	var result *node
	var recur *node
	if a == nil {
		return b
	} else if b == nil {
		return a
	} else {
		recur = ShuffleMergeRecursive(a.next, b.next)
		result = a
		a.next = b
		b.next = recur
		return result
	}
}

func SortedMergeRefs(a, b *node) *node {
	var result *node = nil
	var lastPtrRef **node = &result

	for {
		if a == nil {
			*lastPtrRef = b
			break
		} else if b == nil {
			*lastPtrRef = a
		} else {
			if Less(a.Item, b.Item) {
				MoveNode(lastPtrRef, &a)
			} else {
				MoveNode(lastPtrRef, &b)
			}
			lastPtrRef = &((*lastPtrRef).next)
		}
	}
	return result
}

func SortedMergeRecursive(a, b *node) *node {
	var result *node = nil

	if a == nil {
		return b
	} else if b == nil {
		return a
	}

	if Less(a.Item, b.Item) {
		result = a
		result.next = SortedMergeRecursive(a.next, b)
	} else {
		result = b
		result.next = SortedMergeRecursive(a, b.next)
	}
	return result
}

func SortedMerge(headA, headB *node) *LinkedList {
	var result *node
	last := &result
	for {
		if headA == nil {
			*last = headB
		} else if headB == nil {
			*last = headA
		}

		if Less(headA.Item, headB.Item) {
			MoveNode(last, &headA)
		} else {
			MoveNode(last, &headB)
		}
		last = &((*last).next)
	}
}

func MergeSort(headRef **node) {
	var head *node = *headRef
	var a *node
	var b *node

	if (head == nil) || (head.next == nil) {
		return
	}
	FrontBackSplit(head, &a, &b)
}

func Length(headRef *node) int {
	count := 0
	for curr := headRef; curr != nil; curr = curr.next {
		count++
	}
	return count
}

func FrontBackSplit(source *node, front, back **node) {

	length := Length(source)
	var current *node = source

	if length < 2 {
		*front = source
		*back = nil
	} else {
		hopCount := (length - 1) / 2
		for i := 0; i < hopCount; i++ {
			current = current.next
		}

		*front = source
		*back = current.next
		current.next = nil
	}
}

func FrontBackSplitSlowFast(source *node, front, back **node) {
	var fast *node
	var slow *node
	if source == nil || source.next == nil {
		*front = source
		*back = nil
	} else {
		slow = source
		fast = source.next

		for fast != nil {
			fast = fast.next
			if fast != nil {
				slow = slow.next
				fast = fast.next
			}
		}

		*front = source
		*back = slow.next
		slow.next = nil

	}
}

func (l *LinkedList) DeleteAfter(k int) bool {
	var result bool
	if k > l.Size() {
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

func Less(x, y interface{}) bool {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		panic("mismatched input type")
	}
	switch x.(type) {
	case int:
		a, b := x.(int), y.(int)
		if a > b {
			return false
		}
	case string:
		a, b := x.(string), y.(string)
		if a > b {
			return false
		}
	case float64:
		a, b := x.(float64), y.(float64)
		if a > b {
			return false
		}
	default:
		panic("unhandled type")
	}
	return true
}
