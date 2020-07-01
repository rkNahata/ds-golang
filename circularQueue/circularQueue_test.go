package circularQueue

import "testing"

func TestCircularQueueAll(t *testing.T) {
	q := Init(6)
	if q.size != 6 {
		t.Errorf("expected size to be %d", 6)
	}
	if !q.Enque(1) {
		t.Errorf("expected item to be inserted successfully")
	}
	q.Enque(2)
	_, ok := q.Deque()
	if !ok {
		t.Errorf("expected item to be deleted successfully")
	}

	for i := 1; i <= 6; i++ {
		q.Enque(i * 10)
	}
	q.Display()
}
