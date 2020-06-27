package stack

import "testing"

func TestStackQueueAll(t *testing.T) {
	s := New()
	for i := 0; i < 10; i++ {
		s.Push(i + 1)
	}
	if s.Size() < 10 {
		t.Errorf("expected size 10, got %d", s.Size())
	}
	val := s.Pop()
	if val != 10 {
		t.Errorf("expected value to be 10, got %d", val)
	}
	val = s.Top()
	if val != 9 {
		t.Errorf("expected value to be 9, got %d", val)
	}

}

func BenchmarkStackQueue(b *testing.B) {
	s := New()
	for n := 0; n < b.N; n++ {
		s.Push(10)
	}
}
