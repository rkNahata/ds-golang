package stack

import (
	"testing"
)

func TestQueueStackAll(t *testing.T) {
	q := Initialize()
	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	item := q.Peek()
	if item != 0 {
		t.Error("expected head to be empty")
	}
}
