package bTree

import "ds-golang/queue"

type NodeWithHorizontalDistance struct {
	Left *NodeWithHorizontalDistance
	Right *NodeWithHorizontalDistance
	Val int
	Distance int
}

func LevelOrderTraversal(root *NodeWithHorizontalDistance){
	if root == nil{
		return
	}
	var que = new(queue.Queue)
	horizontalDist := make(map[int][]int)
	horizontalDist[0] = append(horizontalDist[0],root.Val)
	root.Distance = 0
	que.Enqueue(root)
	for que.Size()>0{
		node := que.DequeueRet().(NodeWithHorizontalDistance)
		if node.Left!=nil{
			node.Left.Distance = node.Distance-1
			horizontalDist[node.Left.Distance] = append(horizontalDist[node.Left.Distance],node.Left.Val)
			que.Enqueue(node.Left)
		}
		if node.Right !=nil{
			node.Right.Distance = node.Distance+1
			horizontalDist[node.Right.Distance] = append(horizontalDist[node.Right.Distance],node.Right.Val)
			que.Enqueue(node.Right)
		}
	}

}