package heapsort

import (
	"math/rand"
	"testing"
	"time"
)

const (
	nums = 5
)

func initSorterNums(n int) *Sorter {
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(n)
	t := make([]interface{}, len(nums))
	for i, v := range nums {
		t[i] = v
	}
	s := NewSorter(t)
	return s
}

func TestSorter(t *testing.T) {
	s := initSorterNums(nums)
	expected := nums + 1
	if len(s.Q.pq) != expected {
		t.Error("length of input does not match internal length of sorter interface.")
	}
}

func TestSorted(t *testing.T) {
	s := initSorterNums(nums)
	s.Sort(s.Q.pq)
	s.Show(s.Q.pq)
	if !s.IsSorted(s.Q.pq) {
		t.Error("items not sorted")
	}
}

func TestSortedStructVal(t *testing.T) {
	s := initSorterNums(nums)
	s.Sort(s.Q.pq)
	if s.sorted && !s.IsSorted(s.Q.pq) {
		t.Error("inherent sorted property incorrect.")
	}

}
