package bTree

type Node struct {
	Left  *Node
	Right *Node
	Val   int64
}

func New(val int64) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Val:   val,
	}
}

type Tree struct {
	root *Node
}
