package heapsort

import (
	"fmt"
	"reflect"
)

type MaxPQ struct {
	pq []interface{}
	n  int
}

func NewMaxPQ(maxN int) *MaxPQ {
	var q *MaxPQ = new(MaxPQ)
	q.pq = make([]interface{}, maxN)
	q.n = maxN
	return q
}

func (q *MaxPQ) Size() int {
	return q.n
}

func (q *MaxPQ) IsEmpty() bool {
	return q.n == 0
}

//func (q *MaxPQ) DeleteMax() interface{} {
//	max := q.pq[1]
//	q.Exchange(1, q.n)
//	q.n--
//	lastIndex := q.n + 1
//	q.pq[lastIndex] = nil
//	Sink(1)
//	return max
//
//}

func (q *MaxPQ) Insert(v interface{}) {
	q.n++
	q.pq[q.n] = v
	q.swim(q.n)
}

func (q *MaxPQ) Exchange(i, j int) {
	tmp := q.pq[i-1]
	q.pq[i-1] = q.pq[j-1]
	q.pq[j-1] = tmp
}

func Exchange(input []interface{}, i, j int) {
	tmp := input[i-1]
	input[i-1] = input[j-1]
	input[j-1] = tmp
}

/*
elper function that returns True if j < k.
Off-by-one indexes.
*/
func Less(input []interface{}, j, k int) bool {
	x, y := input[j-1], input[k-1]
	if j-1 == 1 || k-1 == 1 {
		fmt.Println(input)
		fmt.Println("indexes: ", j-1, k-1)
		fmt.Println("vals: ", x, y)
	}
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		panic("mismatched input types")
	}
	switch x.(type) {
	case int:
		a, b := x.(int), y.(int)
		if a < b {
			return true
		}
	case float32:
		a, b := x.(float32), y.(float32)
		if a < b {
			return true
		}
	case float64:
		a, b := x.(float64), y.(float64)
		if a < b {
			return true
		}
	case string:
		a, b := x.(string), y.(string)
		if a < b {
			return true
		}
	default:
		panic("unhandled types, please add case.")
	}
	return false
}

func Sink(input []interface{}, k, n int) {
	for 2*k <= n {
		j := 2 * k
		if j < n && Less(input, j, j+1) {
			j++
		}
		Exchange(input, k, j)
		k = j
	}
}

func (q *MaxPQ) swim(k int) {
	for k > 1 && Less(q.pq, k/2, k) {
		q.Exchange(k/2, k)
		k = k / 2
	}
}
