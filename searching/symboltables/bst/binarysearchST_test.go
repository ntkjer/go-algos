package bst

import (
	"testing"
)

func TestInit(t *testing.T) {
	expected := 1
	bst := New(expected)
	got := bst.n
	if got != expected-1 {
		t.Errorf("expected:[%d] got:[%d]", expected, got)
	}
	got = len(bst.values)
	if got != expected {
		t.Errorf("expected:[%d] got:[%d]", expected, got)
	}
	got = len(bst.keys)
	if got != expected {
		t.Errorf("expected:[%d] got:[%d]", expected, got)
	}
}

func TestPutGetStrings(t *testing.T) {
	size := 1
	length := size + 1 //we are testing two keys in this func
	bst := New(size)
	keyA, keyB := "potato", "avocado"
	valA, valB, valC := "tomato", "toast", "fruit"
	bst.Put(keyA, valA)
	bst.Put(keyB, valB)
	expected := valB
	got := bst.Get(keyB)
	if expected != got {
		t.Error("Error: Get operation on previously inserted key does not work.\n")
		t.Errorf("expected: [%v], got: [%v]\n", expected, got)
	}
	//Put new values on prev key
	bst.Put(keyB, valC)
	got = bst.Get(keyB)
	expected = valC
	if expected != got {
		t.Error("Error: Put operation does not update previously defined k : v0 -> k : v1.\n")
		t.Errorf("expected: [%v], got: [%v]\n", expected, got)
	}
	if length != bst.n {
		t.Error("bst.n incorrect for num items inserted. ")
	}
}

func TestDelete(t *testing.T) {

}

func TestSize(t *testing.T) {

}
