package queue

import (
	"fmt"
	"testing"
)

func TestInitQueue(t *testing.T) {
	var q *Queue = new(Queue)
	if q.Size() != 0 {
		t.Error("expected 0, got ", q.Size())
	}
}

func TestEnqeueThenDequeue(t *testing.T) {
	var q *Queue = new(Queue)

	testItems := []string{
		"apple",
		"banana",
		"carrot",
	}
	for _, item := range testItems {
		q.Enqueue(item)
	}
	if q.Size() != len(testItems) {
		t.Errorf("expected %d got %d \n.", len(testItems), q.Size())
	}
	fmt.Println(testItems)
	fmt.Println(q.first.item)
	for i := 0; i < len(testItems) && !q.IsEmpty(); i++ {
		got := q.Dequeue()
		expected := testItems[i]
		if got != expected {
			t.Errorf("expected: %v but got :%v\n.", expected, got)
		}
	}
	if q.IsEmpty() {
		if q.Size() != 0 {
			t.Error("expeted empty queue to have 0 size.")
		}
	}
}
