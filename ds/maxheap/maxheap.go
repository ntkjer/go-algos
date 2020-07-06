package maxheap

import (
	"log"

	"github.com/ntkjer/sedgewick/ds/heap"
)

type MaxHeap struct {
	*heap.Heap
}

func New(input []int) *MaxHeap {
	h := &MaxHeap{
		&heap.Heap{
			Items: input,
		},
	}

	if len(h.Items) > 0 {
		h.buildMaxHeap()
	}

	return h
}

func (h *MaxHeap) buildMaxHeap() {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.maxHeapifyDown(i)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {

	// iterate until we have achild node smaller than the index value
	for (h.HasLeft(index) && h.Items[index] < h.Left(index)) ||
		(h.HasRight(index) && h.Items[index] < h.Right(index)) {
		// if both children are smaller
		if (h.HasLeft(index) && h.Items[index] < h.Left(index)) &&
			(h.HasRight(index) && h.Items[index] < h.Right(index)) {
			//Compare children and swap the largest item
			if h.Left(index) > h.Right(index) {
				h.Swap(index, h.GetLeftIndex(index))
				index = h.GetLeftIndex(index)
			} else {
				h.Swap(index, h.GetRightIndex(index))
			}
			// If left child is smaller, then swap it
		} else if h.HasLeft(index) && h.Items[index] < h.Left(index) {
			h.Swap(index, h.GetLeftIndex(index))
			index = h.GetLeftIndex(index)
			// Right child must be smaller, so swap it
		} else {
			h.Swap(index, h.GetRightIndex(index))
			index = h.GetRightIndex(index)
		}
	}
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.HasParent(index) && h.Parent(index) < h.Items[index] {
		h.Swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
	}
}

func (h *MaxHeap) Insert(item int) *MaxHeap {
	h.Items = append(h.Items, item)
	lastElementIndex := len(h.Items) - 1
	h.maxHeapifyUp(lastElementIndex)
	return h
}

func (h *MaxHeap) DeleteMax() int {
	if len(h.Items) == 0 {
		log.Fatal("no items in heap.")
	}

	minItem := h.Items[0]
	lastIndex := len(h.Items) - 1
	h.Items[0] = h.Items[lastIndex]

	h.Items = h.Items[:len(h.Items)-1]

	h.maxHeapifyDown(0)

	return minItem
}
