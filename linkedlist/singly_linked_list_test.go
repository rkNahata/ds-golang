package linkedlist

import (
	"testing"
)

func TestSinglyLinkedListAll(t *testing.T) {
	var ll SinglyLinkedList
	var count = 7

	if !ll.IsEmpty() {
		t.Error("expected list to be empty")
	}

	ll.BulkInsert(count)

	if ll.IsEmpty() {
		t.Error("expected list to be not empty")
	}
	if ll.Length() != 7 {
		t.Errorf("length expected %d but it is %d", count, ll.Length())
	}

	if err := ll.InsertAt(count-1, 10); err != nil {
		t.Error("failed to insert item", err)
	}
	if err := ll.InsertAt(count+2, 10); err == nil {
		t.Error("expected index error")

	}
	if err := ll.RemoveAt(count); err != nil {
		t.Error("failed to insert item", err)
	}
	if err := ll.RemoveAt(count + 1); err == nil {
		t.Error("expected index error")
	}
	if ll.Length() != count {
		t.Errorf("expected lenght %d got %d", count, ll.Length())
	}

	found, index := ll.IndexOf(4)
	if !found {
		t.Error("expected element to be found")
	}
	if index != 4 {
		t.Errorf("expected index of the element to be %d got %d", 4, index)
	}
	ll.ShowAll()

}
