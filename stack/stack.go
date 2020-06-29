package stack

type Stacker interface {
	Push()
	Pop()
	Peek()
}

type element struct {
	value interface{}
	next  *element
}

type Stack struct {
	top  *element
	size int
}

func New() *Stack {
	return &Stack{top: nil, size: 0,}
}

func (s *Stack) Push(val interface{}) {
	s.top = &element{value: val, next: s.top,}
	s.size++
}

func (s *Stack) Pop() (val interface{}, exists bool) {
	if s.size > 0 {
		val = s.top.value
		s.top = s.top.next
		s.size--
		exists = true
	}
	return
}

func (s *Stack) Peek() (val interface{}, exists bool) {
	exists = false
	if s.size > 0 {
		val = s.top.value
		exists = true
	}
	return
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
