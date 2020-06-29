package stack

import "ds-golang/queue"

type StackQueue struct {
	q *queue.Queue
}

func Constructor() *StackQueue {
	q := queue.New()
	return &StackQueue{q}
}

func (sq *StackQueue) Push(item int) {
	sq.q.Enqueue(item)
	size := sq.q.Size()
	for size > 1 {
		sq.q.Enqueue(sq.q.DequeueRet().Value.(int))
		size--
	}
}

func (sq *StackQueue) Pop() int {
	node := sq.q.DequeueRet()
	if node != nil {
		return node.Value.(int)
	}
	return -1
}

func (sq *StackQueue) Top() int {
	if sq.q.Peek() != nil {
		return sq.q.Peek().Value.(int)
	}
	return -1
}

func (sq *StackQueue) Size() int {
	return sq.q.Size()
}
