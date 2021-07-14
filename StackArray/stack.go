package StackArray

import "fmt"

func AllPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	dfs(graph, &result, &path, 0, len(graph)-1)
	return result

}

func dfs(graph [][]int, result *[][]int, path *[]int, vertex, dest int) {
	*path = append(*path, vertex)
	if vertex == dest {
		*result = append(*result, *path)
	} else {
		for _, v := range graph[vertex] {
			dfs(graph, result, path, v, dest)
		}
	}
	temp := *path
	temp = temp[:len(temp)-1]
	path = &temp
}

func Runner() {
	type Tests struct {
		Graph [][]int
	}
	tests := []Tests{ /*{[][]int{{3, 1}, { 2, 4}, {3}, {4}, {}}},*/
		{[][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}},}

	for _, t := range tests {
		result := AllPathsSourceTarget(t.Graph)
		fmt.Println(result)
	}
}

