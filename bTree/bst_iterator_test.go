package bTree

import (
	"fmt"
	"testing"
)

func TestBSTIteratorAll(t *testing.T) {
	var tree Tree
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(15)
	tree.Insert(13)
	tree.Insert(19)

	bst := InitBST(tree.root)
	for bst.HasNext(tree.root) {
		fmt.Println(bst.Next())
	}
}
