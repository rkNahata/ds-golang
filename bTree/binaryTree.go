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

func (n *Node) Insert(val int64) {
	if n.Val > val {
		if n.Left != nil {
			n.Left.Insert(val)
		} else {
			n.Left.Val = val
		}
	} else {
		if n.Right != nil {
			n.Right.Insert(val)
		} else {
			n.Right.Val = val
		}
	}
}

func (n *Node) Find(val int64) bool {
	switch {
	case val == n.Val:
		return true
	case val < n.Val:
		if n.Left != nil {
			return n.Left.Find(val)
		}
	case val > n.Val:
		if n.Right != nil {
			return n.Right.Find(val)
		}
	}
	return false
}
