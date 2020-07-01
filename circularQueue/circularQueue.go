package circularQueue

import "fmt"

type CircularQueue struct {
	front, rear, size int
	array             []int
}

func Init(size int) *CircularQueue {
	return &CircularQueue{
		front: -1,
		rear:  -1,
		size:  size,
		array: make([]int, size),
	}
}

func (cq *CircularQueue) Enque(item int) bool {
	if (cq.front == 0 && cq.rear == cq.size-1) || (cq.rear == (cq.front-1)%(cq.size-1)) {
		fmt.Println("queue is full. No further insertion possible")
		return false
	}
	if cq.front == -1 {
		cq.rear, cq.front = 0, 0
		cq.array[cq.rear] = item
	} else if cq.front != 0 && cq.rear == cq.size-1 {
		cq.rear = 0
		cq.array[cq.rear] = item
	} else {
		cq.rear++
		cq.array[cq.rear] = item
	}
	return true
}

func (cq *CircularQueue) Deque() (int, bool) {
	if cq.front == -1 {
		fmt.Println("queue is empty.deletion not possible")
		return -1, false
	}
	data := cq.array[cq.front]
	cq.array[cq.front] = 0
	if cq.front == cq.rear {
		cq.front, cq.rear = -1, -1
	} else if cq.front == cq.size-1 {
		cq.front = 0
	} else {
		cq.front++
	}
	return data, true
}

func (cq *CircularQueue) Display() {
	if cq.front == -1 {
		fmt.Println("queue is empty")
		return
	}
	if cq.rear >= cq.front {
		for i := cq.front; i <= cq.rear; i++ {
			fmt.Println(cq.array[i])
		}

	} else {
		for i := cq.front; i < cq.size; i++ {
			fmt.Println(cq.array[i])
		}
		for i := 0; i <= cq.rear; i++ {
			fmt.Println(cq.array[i])
		}
	}
}
