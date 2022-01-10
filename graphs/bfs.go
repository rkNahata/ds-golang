package graphs

import (
	"fmt"
	"github.com/rknahata/ds-golang/linkedlist"
	"github.com/rknahata/ds-golang/queue"
)

func traversal(graph [][]int, start int) {
	que := queue.New()
	visited := make([]int, len(graph))
	i := start
	visited[i] = 1
	fmt.Println(i)
	que.Enqueue(i)
	for que.Size() > 0 {
		vNode, _ := que.Peek().(*linkedlist.Node)

		v, _ := vNode.Value.(int)

		que.Dequeue()
		for j := 1; j < len(graph[v]); j++ {
			if graph[v][j] == 1 && visited[j] == 0 {
				fmt.Println(j)
				visited[j] = 1
				que.Enqueue(j)
			}
		}
	}

}

func traversalDFS(graph [][]int, start int, visited []int) {
	if visited[start] == 0 {
		fmt.Print(start)
		visited[start] = 1
		for j := 1; j < len(graph); j++ {
			if graph[start][j] == 1 && visited[j] == 0 {
				traversalDFS(graph, j, visited)
			}
		}
	}
}

func Runner() {
	g := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{0, 0, 1, 1, 0, 1, 1},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},}
	visited := make([]int, len(g))

	traversalDFS(g, 4, visited)

}
