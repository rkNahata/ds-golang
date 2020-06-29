package stack

type QueueStack struct {
	s1, s2 *Stack
	front  int
}

func Initialize() *QueueStack {
	return &QueueStack{
		s1:    New(),
		s2:    New(),
		front: -1,
	}
}

func (qs *QueueStack) IsEmpty() bool {
	return qs.s1.IsEmpty() && qs.s2.IsEmpty()
}

func (qs *QueueStack) Push(item int) {
	if qs.s1.IsEmpty() {
		qs.front = item
	}
	qs.s1.Push(item)
}

func (qs *QueueStack) Pop() int {
	if qs.s2.IsEmpty() {
		for !qs.s1.IsEmpty() {
			item, _ := qs.s1.Pop()
			qs.s2.Push(item)
		}
	}
	pop, _ := qs.s2.Pop()
	return pop.(int)
}

func (qs *QueueStack) Peek() int {
	if !qs.s2.IsEmpty() {
		top, _ := qs.s2.Peek()
		return top.(int)
	}
	return qs.front

}
