package dfs

import "fmt"

func CanFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	graph := make(map[int][]int)
	for i := range prerequisites {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}

	visited := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		if isCycle(graph, &visited, i) {
			return false
		}
	}
	return true
}

func isCycle(graph map[int][]int, visited *[]int, node int) bool {
	v := *visited
	if v[node] == 1 {
		return true
	}
	if v[node] == 0 {
		v[node] = 1
		for i, _ := range graph[node] {
			if isCycle(graph, &v, graph[node][i]) {
				return true
			}
		}
	}
	v[node] = 2
	visited = &v
	return false
}

func Runner() {
	//preRequisites := [][]int{{0, 1}, {1, 0}}
	preRequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	numCourses := 4
	fmt.Println(CanFinish(numCourses, preRequisites))
}
