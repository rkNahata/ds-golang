package dll

import (
	"fmt"
	"testing"
)

func TestDoublyLinkedListAll(t *testing.T) {
	ll := New()

	for i := 0; i < 10; i++ {
		ll.Append(i+1)
	}
	if ll.Size() != 10 {
		t.Errorf("size expected 10 but it is %d", ll.Size())
	}

	ll.Prepend(11)
	if ll.Size() != 11 {
		t.Errorf("size expected 11 but it is %d", ll.Size())
	}

	ll.DeleteLast()
	if ll.Size() != 10 {
		t.Errorf("size expected 10 but it is %d", ll.Size())
	}

	ll.Prepend(12)
	for ll.head!=nil{
		ll.DeleteLast()
	}
	fmt.Println(ll.LookUp())

}
