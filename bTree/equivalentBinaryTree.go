package bTree

import (
	"fmt"
	"math/rand"
)

/*A function to check whether two binary trees store the same sequence*/

func Create(k int) *Node {
	var t *Node
	for _, v := range rand.Perm(10) {
		t = insert(t, int64((1+v)*k))
	}
	return t
}

func insert(t *Node, v int64) *Node {
	if t == nil {
		return &Node{nil, nil, v}
	}
	if v < t.Val {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func Same(t1, t2 *Node) bool {
	c1, c2 := make(chan int64), make(chan int64)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		x1, ok1 := <-c1
		x2, ok2 := <-c2
		switch {
		//not same size
		case ok1 != ok2:
			return false
		//not same value
		case x1 != x2:
			return false
		case !ok1:
			return true
		default:
			//keep iterating
		}
	}

}

func Walk(t *Node, c chan int64) {
	recWalk(t, c)
	close(c)
}

func recWalk(t *Node, c chan int64) {
	if t != nil {
		recWalk(t.Left, c)
		c <- t.Val
		recWalk(t.Right, c)
	}
}

func ExecuterCode() {
	t1 := Create(1)
	t2 := Create(2)
	c := make(chan int64)
	go Walk(t1, c)
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println(Same(t1, t1))
	fmt.Println(Same(t1, t2))
}
