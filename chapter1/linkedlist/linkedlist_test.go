package linkedlist

import "testing"

func TestInitSize(t *testing.T) {
	var l *LinkedList = new(LinkedList)
	if l.Size() != 0 {
		t.Error("expected 0, got ", l.Size())
	}
}

func TestSingleItem(t *testing.T) {
	var l *LinkedList = new(LinkedList)
	l.InsertFirst(1)
	if l.Size() != 1 {
		t.Error("expected 1, got ", l.Size())
	}
	if l.first.Item != 1 {
		t.Errorf("expected first 1 got %v\n", l.first.Item)
	}
	if l.first.next != nil {
		t.Error("expected first val to have next item of nil, got ", l.first.next)
	}
	l.DeleteFirst()
	if l.Size() != 0 {
		t.Error("expected a size of zero after deleting first item.")
	}
	if l.first != nil {
		t.Error("expected first val to be nil, got ", l.first.Item)
	}
}

func TestMultipleItems(t *testing.T) {
	var l *LinkedList = new(LinkedList)
	l.InsertFirst(1)
	l.InsertFirst(0)
	l.Append(2)

	if l.first.Item != 0 {
		t.Error("expected 0, got ", l.first)
	}

	expected := l.Find(1)
	if expected != 1 {
		t.Error("expected 1st pos for int 1, got ", expected)
	}
	expected = l.Find(2)
	if expected != 2 {
		t.Error("expected 2nd pos for int 2, got ", expected)
	}

	expected = l.max(l.first)
	if expected != 2 {
		t.Error("expected max to be 2 got ", expected)
	}

	expected = l.RecursiveMax(l.first, 0)
	if expected != 2 {
		t.Error("expected max to be 2 got ", expected)
	}

	expected = 2
	got := Reverse(l.first).Item
	if got != expected {
		t.Errorf("expected reversed item was %d , got %d \n", expected, got)
	}

	expected = 0
	got = rReverse(l.first).Item
	if got != expected {
		t.Errorf("expected reversed item was %d , got %d \n", expected, got)
	}

	l.DeleteLast()
	got = l.Find(2)
	expected = -1
	if got != expected {
		t.Errorf("expected reversed item was %d , got %d \n", expected, got)
	}

}

//		item := scanner.Text()
//		if item == "" {
//			break
//		} else if item == "dl" {
//			fmt.Println("calling delete last, Size is: ", l.Size())
//			l.DeleteLast()
//			fmt.Println("called delete last, Size is now: ", l.Size())
//			fmt.Println("head is at -> ", l.first)
//		} else if item == "." {
//			fmt.Printf("deleted first: %v <-- length: %v\n", l.DeleteFirst(), l.Size())
//		} else if item == "da" {
//			fmt.Println(l.DeleteAfter(l.Size()/2 - 1))
//			l.PrintLinks()
//		} else if item == "Find" {
//			fmt.Println(l.Find("niels"))
//		} else if item == "ia" {
//			i := 3
//			l.AddAfter(i, "baboom")
//		} else if item == "pl" {
//			l.PrintLinks()
//		} else if item == "max" {
//			fmt.Println(l.max(l.first))
//		} else if item == "rmax" {
//			max := 0
//			fmt.Println(l.RecursiveMax(l.first, max))
//		} else if item == "Reverse" {
//			l.first = Reverse(l.first)
//		} else if item == "rReverse" {
//			l.first = rReverse(l.first)
//		} else {
//			num, _ := strconv.Atoi(item)
//			l.InsertFirst(num)
//			//l.InsertFirst(item)
//		}
//
