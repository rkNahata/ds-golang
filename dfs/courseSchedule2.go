package dfs

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	topologicalOrder := make([]int, numCourses)
	inDegree := make([]int, numCourses)
	stk := New()
	for i := range prerequisites {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
		inDegree[prerequisites[i][0]]++

	}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			stk.Push(inDegree[i])
		}
	}
	var i int
	for !stk.IsEmpty() {
		node := stk.Pop()
		topologicalOrder[i] = node
		i++

		for _, neighbour := range graph[node] {
			inDegree[neighbour]--
			if inDegree[neighbour] == 0 {
				stk.Push(neighbour)
			}
		}
	}
	if i == numCourses {
		return topologicalOrder
	}
	return []int{0}
}

func Execute() {
	preRequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	numCourses := 4
	fmt.Println(findOrder(numCourses, preRequisites))
}
