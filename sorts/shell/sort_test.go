package shell

import (
	"math/rand"
	"testing"
	"time"
)

func initSorterNums(n int) *Sorter {
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(n)
	var s *Sorter = new(Sorter)
	t := make([]interface{}, len(nums))
	for i, v := range nums {
		t[i] = v
	}
	s.items = t
	s.sorted = false
	return s
}

func TestSorter(t *testing.T) {
	nums := 1000
	s := initSorterNums(nums)
	if len(s.items) != nums {
		t.Error("length of input does not match internal length of sorter interface.")
	}
}

func TestSorted(t *testing.T) {
	nums := 1000
	s := initSorterNums(nums)
	s.Sort(s.items)
	if !s.IsSorted(s.items) {
		t.Error("items not sorted")
	}
}

func TestSortedStructVal(t *testing.T) {
	nums := 10000
	s := initSorterNums(nums)
	s.Sort(s.items)
	if s.sorted && !s.IsSorted(s.items) {
		t.Error("inherent sorted property incorrect.")
	}

}
