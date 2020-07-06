package minheap

import (
	"log"
	"reflect"

	"github.com/ntkjer/sedgewick/ds/heap"
)

type MinHeap struct {
	*heap.Heap
}

func New(input []interface{}) *MinHeap {
	h := &MinHeap{
		&heap.Heap{
			Items: input,
		},
	}

	if len(h.Items) > 0 {
		h.buildMinHeap()
	}

	return h
}

func (h *MinHeap) buildMinHeap() {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.minHeapifyDown(i)
	}
}

func (h *MinHeap) minHeapifyDown(index int) {

	// iterate until we have achild node smaller than the index value
	for (h.HasLeft(index) && h.Less(h.Left(index), h.Items[index])) ||
		(h.HasRight(index) && h.Less(h.Right(index), h.Items[index])) {
		// if both children are smaller
		if (h.HasLeft(index) && h.Less(h.Left(index), h.Items[index])) &&
			(h.HasRight(index) && h.Less(h.Right(index), h.Items[index])) {
			//Compare children and swap the largest item
			if h.Less(h.Left(index), h.Right(index)) {
				h.Swap(index, h.GetLeftIndex(index))
				index = h.GetLeftIndex(index)
			} else {
				h.Swap(index, h.GetRightIndex(index))
			}
			// If left child is smaller, then swap it
		} else if h.HasLeft(index) && h.Less(h.Left(index), h.Items[index]) {
			h.Swap(index, h.GetLeftIndex(index))
			index = h.GetLeftIndex(index)
			// Right child must be smaller, so swap it
		} else {
			h.Swap(index, h.GetRightIndex(index))
			index = h.GetRightIndex(index)
		}
	}
}

func (h *MinHeap) Less(x, y interface{}) bool {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		panic("mismatched input types for args x and y.")
	}

	switch x.(type) {
	case float64:
		a, b := x.(float64), y.(float64)
		if a > b {
			return false
		}
	default:
		panic("unhandled type.")
	}
	return true
}
func (h *MinHeap) minHeapifyUp(index int) {
	for h.HasParent(index) && h.Less(h.Items[index], h.Parent(index)) {
		h.Swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
	}
}

func (h *MinHeap) Insert(item interface{}) *MinHeap {
	h.Items = append(h.Items, item)
	lastElementIndex := len(h.Items) - 1
	h.minHeapifyUp(lastElementIndex)
	return h
}

func (h *MinHeap) DeleteMin() interface{} {
	if len(h.Items) == 0 {
		log.Fatal("no items in heap.")
	}

	minItem := h.Items[0]
	lastIndex := len(h.Items) - 1
	h.Items[0] = h.Items[lastIndex]

	h.Items = h.Items[:len(h.Items)-1]

	h.minHeapifyDown(0)

	return minItem
}
