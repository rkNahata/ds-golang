package bTree

type BSTIterator struct {
	stack []*Node
}

func InitBST(root *Node) *BSTIterator {
	var stack []*Node
	bst := &BSTIterator{stack: stack}
	bst.LeftmostInorder(root)
	return bst
}

func (bst *BSTIterator) LeftmostInorder(root *Node) {
	for root != nil {
		bst.stack = append(bst.stack, root)
		root = root.Left
	}
}

func (bst *BSTIterator) Next() int64 {
	n := bst.stack[len(bst.stack)-1]
	bst.stack = bst.stack[:len(bst.stack)-1]
	if n.Right != nil {
		bst.LeftmostInorder(n.Right)
	}
	return n.Val
}

func (bst *BSTIterator) HasNext(root *Node) bool {
	return len(bst.stack) != 0
}
