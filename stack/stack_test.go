package stack

import (
	"fmt"
	"testing"
)

func TestStackAll(t *testing.T) {
	s := new(Stack)
	for i := 0; i < 10; i++ {
		s.Push(i + 1)
	}
	if s.size < 10 {
		t.Errorf("expected size 10, got %d", s.size)
	}
	val, exist := s.Pop()
	if !exist {
		t.Error("expected exist to be true, got false")
	}
	if val != 10 {
		t.Errorf("expected value to be 10, got %d", val)
	}
	val, exist = s.Peek()
	if !exist {
		t.Error("expected exist to be true, got false")
	}
	if val != 9 {
		t.Errorf("expected value to be 10, got %d", val)
	}

}

func TestMinStack(t *testing.T) {
	ms := new(MinStack)
	values := []int{2, 6, 5, 3, 1}

	for i := range values {
		ms.Push(values[i])
	}
	fmt.Printf("%v", ms.size)
	fmt.Printf("%v", ms.min)
	if ms.size != len(values) {
		t.Errorf("expected size %d, got %d", len(values), ms.size)
	}
	if ms.min != 1 {
		t.Errorf("expected min to be %d, got %d", 1, ms.min)
	}
	ms.Pop()
	ms.Pop()
	ms.Pop()
	ms.Pop()
	if ms.size != len(values)-4 {
		t.Errorf("expected size %d, got %d", len(values)-4, ms.size)
	}
	if ms.min != values[0] {
		t.Errorf("expected min to be %d, got %d", values[0], ms.min)
	}
	ms.Pop()
	_, exist := ms.Pop()
	if exist {
		t.Error("expected exist to be false, got true")
	}
}


func BenchmarkStack(b *testing.B) {
	s := new(Stack)
	for n := 0; n < b.N; n++ {
		s.Push(10)
	}
}