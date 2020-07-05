package bTree

func (t *Tree) Insert(data int64) {
	if t.root == nil {
		t.root = &Node{Left: nil, Right: nil, Val: data,}
	} else {
		t.root.insert(data)

	}
}

func (n *Node) insert(data int64) {
	if n.Val > data {
		if n.Left == nil {
			n.Left = New(data)
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = New(data)
		} else {
			n.Right.insert(data)
		}
	}
}
