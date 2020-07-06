package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Heap struct {
	PQ []interface{}
}

func (h *Heap) Sort(pq []interface{}) {
	n := len(pq)

	for k := n / 2; k >= 1; k-- {
		Sink(pq, k, n)
	}
	for n > 1 {
		Exchange(pq, 1, n)
		n = n - 1
		//Sink(pq, 1, n)
		SiftDown(pq, 1, n, 1)
	}
}

func SiftDown(pq []interface{}, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && Less(pq, first+child, first+child+1) {
			child++
		}

		if !Less(pq, first+root, first+child) {
			return
		}
		Exchange(pq, first+root, first+child)
		root = child
	}
}

func Sink(pq []interface{}, k, n int) {
	for 2*k <= n {
		j := 2 * k
		fmt.Println("sink state: ", j, k)
		if j < n && Less(pq, j, j+1) {
			j = j + 1
		}
		Exchange(pq, k, j)
		k = j
	}
}

func Exchange(pq []interface{}, j, k int) {
	curr := pq[j-1]
	pq[j-1] = pq[k-1]
	pq[k-1] = curr
	curr = nil
}

func Less(pq []interface{}, j, k int) bool {
	x, y := pq[j-1], pq[k-1]
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		fmt.Println("mismatched inputs", x, y)
		panic("mismatched inputs")
	}
	switch x.(type) {
	case int:
		a, b := x.(int), y.(int)
		if a > b {
			return false
		}
	case float32:
		a, b := x.(int), y.(int)
		if a > b {
			return false
		}
	case float64:
		a, b := x.(int), y.(int)
		if a > b {
			return false
		}
	case string:
		a, b := x.(string), y.(string)
		if a > b {
			return false
		}
	default:
		panic("unhandled types, please add case.")
	}
	return true
}

func main() {
	a := readStdin()
	var h *Heap = new(Heap)
	h.PQ = a
	h.Sort(h.PQ)
	fmt.Println(h.PQ)
}

func readStdin() []interface{} {
	scanner := bufio.NewScanner(os.Stdin)
	var items []interface{}
	for scanner.Scan() {
		item := scanner.Text()
		tmp := strings.SplitAfter(item, " ")
		items = make([]interface{}, len(tmp)+1)
		for i, item := range tmp {
			items[i+1] = item
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return items
}
