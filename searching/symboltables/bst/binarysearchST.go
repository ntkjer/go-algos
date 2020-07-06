package bst

import (
	"reflect"
)

type BinarySearchST struct {
	keys   []interface{}
	values []interface{}
	n      int
}

func New(capacity int) *BinarySearchST {
	keys := make([]interface{}, capacity)
	vals := make([]interface{}, capacity)
	var b *BinarySearchST = new(BinarySearchST)
	b.keys = keys
	b.values = vals
	b.n = 0
	return b
}

func (b *BinarySearchST) Size() int {
	return b.n
}

func (b *BinarySearchST) isEmpty() bool {
	return b.Size() == 0
}

func (b *BinarySearchST) Get(key interface{}) interface{} {
	if b.isEmpty() {
		return nil
	}
	i := b.rank(key)
	if i < b.n && equals(b.keys[i], key) {
		return b.values[i]
	}
	return nil
}

func (b *BinarySearchST) Put(key, value interface{}) {
	i := b.rank(key)
	if i < b.n && compareTo(b.keys[i], key) == 0 {
		b.values[i] = value
		return
	}
	for j := b.n; j > i; j-- {
		b.keys[j] = b.keys[j-1]
		b.values[j] = b.values[j-1]
	}
	if b.n == i {
		b.resize()
	}
	b.keys[i] = key
	b.values[i] = value
	b.n++
}

func (b *BinarySearchST) resize() {
	resizeFactor := b.n*2 + 1

	newKeys := make([]interface{}, resizeFactor)
	newVals := make([]interface{}, resizeFactor)
	newKeys = append(newKeys, b.keys[:b.n])
	newVals = append(newVals, b.values[:b.n])

	//	for i := 0; i < b.n; i++ {
	//		newKeys[i] = b.keys[i]
	//		newVals[i] = b.values[i]
	//	}

	b.keys, b.values = newKeys, newVals
}

func (b *BinarySearchST) rank(key interface{}) int {
	lo := 0
	hi := b.Size() - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		cmp := compareTo(key, b.keys[mid])
		if cmp > 0 {
			lo = mid + 1
		} else if cmp < 0 {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return lo
}

func (b *BinarySearchST) Delete(key interface{}) {
	if b.isEmpty() {
		return
	}
	i := b.rank(key)
	if i < b.n && equals(b.keys[i], key) {
		b.keys[i] = nil
		b.values[i] = nil
	}
}

// Returns -1 if less, 0 if equal, 1 if x greater than y.
func compareTo(x, y interface{}) int {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		if reflect.TypeOf(x) == nil {
			return -1
		} else if reflect.TypeOf(x) == reflect.TypeOf(y) {
			return 0
		} else {
			return 1
		}
	}
	var result int
	switch x.(type) {
	case int:
		a, b := x.(int), y.(int)
		if a > b {
			result = 1
		} else if a == b {
			result = 0
		} else {
			result = -1
		}
	case float64:
		a, b := x.(float64), y.(float64)
		if a > b {
			result = 1
		} else if a == b {
			result = 0
		} else {
			result = -1
		}
	case string:
		a, b := x.(string), y.(string)
		if a > b {
			result = 1
		} else if a == b {
			result = 0
		} else {
			result = -1
		}
	default:
		panic("add new case")
	}
	return result
}

func equals(x, y interface{}) bool {
	return compareTo(x, y) == 0
}
