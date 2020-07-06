package heap

import "fmt"

type Heap struct {
}

func heapSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	//Build heap with greatest elements at the top
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}
	//pop elements, largest first, at the end of the data.
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+1)
		siftDown(data, lo, i, first)
	}
}

func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

// A collection type that can satisfy the underlying interface of the Heap for varuous types.
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int) bool
}

// Type Wrappers for various interface : type mappings

/*
Strings
*/
type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p StringSlice) Sort() { Sort(p) }

/*
Ints
*/
type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p IntSlice) Sort() { Sort(p) }

/*
Float64s
*/
type Float64Slice []float64

func (p Float64Slice) Len() int {
	return len(p)
}

func (p Float64Slice) Less(i, j int) bool {
	return p[i] < p[j] || isNan(p[i]) && !isNan(p[j])
}

func (p Float64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func isNan(f float64) bool {
	return f != f
}

func (p Float64Slice) Sort() { Sort(p) }

// Embedded interfaces
type reverse struct {
	// This embedded Interface permits REverse to use the methods of other interfaces.
	Interface
}

func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

// Interface methods and utils
func Reverse(data Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

func Sort(data Interface) {
	n := data.Len()
	fmt.Println(n)
	//Sort here
}

func IsSorted(data Interface) bool {

}

// Convenience wrappers
func Strings(a []string) {
	Sort(StringSlice(a))
}
func Ints(a []int) {
	Sort(IntSlice(a))
}

func Float64s(a []float64) {
	Sort(Float64Slice(a))
}
