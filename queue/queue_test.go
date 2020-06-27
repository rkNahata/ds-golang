package queue

import (
	"testing"
)

func TestQueueAll(t *testing.T) {
	q := New()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	if q.size < 10 {
		t.Errorf("expected size 10, got %d", q.size)
	}

	for i := 10; i > 0; i-- {
		q.Dequeue()
	}

	if q.size != 0 {
		t.Error("expected queue to be empty")
	}

	node := q.Peek()
	if node != nil {
		t.Error("expected head to be empty")
	}
}
