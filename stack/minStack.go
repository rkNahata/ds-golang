package stack

type MinStack struct {
	top  *element
	size int
	min  int
}

func (ms *MinStack) Push(val interface{}) {
	var temp int
	if ms.size == 0 {
		ms.top = &element{value: val, next: ms.top}
		ms.min = val.(int)
	} else {
		if ms.min > val.(int) {
			temp = 2*(val.(int)) - ms.min
			ms.min = val.(int)
		} else {
			temp = val.(int)
		}
		ms.top = &element{value: temp, next: ms.top}
	}
	ms.size++
}

func (ms *MinStack) Pop() (val interface{}, exist bool) {
	exist = false
	if ms.size == 0 {
		return
	}
	exist = true
	if ms.min > ms.top.value.(int) {
		tmp := 2*ms.min - ms.top.value.(int)
		val = ms.min
		ms.min = tmp

	} else {
		val = ms.top.value
	}
	ms.top = ms.top.next
	ms.size--
	return
}
