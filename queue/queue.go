package queue

import (
	"github.com/rknahata/ds-golang/linkedlist"
)

type Queue struct {
	list *linkedlist.SinglyLinkedList
	size int
}

func New() *Queue {
	list := linkedlist.New()
	return &Queue{list, 0}
}

func (q *Queue) Enqueue(item interface{}) {
	q.list.Append(item)
	q.size++
}

func (q *Queue) Dequeue() {
	if q.size > 0 {
		q.list.DeleteAtTop()
		q.size--
	}
}

func (q *Queue)DequeueRet()interface{}{
	node := q.Peek()
	if node !=nil{
		q.list.DeleteAtTop()
		q.size--
	}
	return node
}

func (q *Queue) Peek() interface{} {
	return q.list.Head
}

func (q *Queue)Size()int{
	return q.size
}

func (q *Queue)DequeueTreeNode()interface{}{
	node := q.Peek()
	if node !=nil{
		q.list.DeleteAtTop()
		q.size--
	}
	return node
}