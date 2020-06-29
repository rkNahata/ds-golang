package stack

import "testing"

func TestStackQueueAll(t *testing.T) {
	s := New()
	for i := 0; i < 10; i++ {
		s.Push(i + 1)
	}
}
