package interleaver

import (
	"fmt"
	"github.com/rknahata/ds-golang/linkedlist"
	"github.com/rknahata/ds-golang/queue"
	"log"
)

type Interleaver struct {
	array  [][]int
	cycle  bool
	Queue  *queue.Queue
	result []int
}

func NewInterleaver(array [][]int, cycle bool) *Interleaver {
	qu := queue.New()
	return &Interleaver{
		array: array,
		cycle: cycle,
		Queue: qu,
	}
}

func (in *Interleaver) Next() int {
	if in.Queue.Size() == 0 {
		if in.cycle {
			in.Reset()
		} else {
			log.Fatalln("something went wrong")
		}
	}
	eleInter := in.Queue.Peek().(*linkedlist.Node)
	in.Queue.Dequeue()
	elements, _ := eleInter.Value.([]int)

	i := elements[0]
	j := elements[1]
	if j < len(in.array[i])-1 {
		in.Queue.Enqueue([]int{i, j + 1})
	}
	in.result = append(in.result, in.array[i][j])
	return in.array[i][j]
}

func (in *Interleaver) Reset() {
	for i := range in.array {
		if len(in.array[i]) > 0 {
			in.Queue.Enqueue([]int{i, 0})
		}
	}
}

func (in *Interleaver) HasNext() bool {
	return in.cycle || in.Queue.Size() > 0
}

func Runner() {
	input := [][]int{{1, 2, 3}, {4, 5}, {6}, {}, {7, 8, 9}}

	//[1,4,6,7,2,5,8,3,9]
	runner := NewInterleaver(input, true)
	for i := 0; i < 9; i++ {
		fmt.Println(runner.Next())
	}
	fmt.Println(runner.result)
}
