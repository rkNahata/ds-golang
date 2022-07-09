package graphs

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	n := len(heights)
	m := len(heights[0])
	altVisited := make([][]bool, n)
	pacVisited := make([][]bool, n)

	initVisited(altVisited)
	initVisited(pacVisited)
	var pacQueue, altQueue Queue

	for i := 0; i < n; i++ {
		pacQueue.Enqueue([]int{i, 0})
		altQueue.Enqueue([]int{i, m - 1})
		pacVisited[i][0] = true
		altVisited[i][m-1] = true

	}

	for i := 0; i < m; i++ {
		pacQueue.Enqueue([]int{0, i})
		altQueue.Enqueue([]int{n - 1, i})
		pacVisited[0][i] = true
		altVisited[n-1][i] = true

	}

	bfs(heights, &pacQueue, pacVisited)
	bfs(heights, &altQueue, altVisited)
	var result [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if pacVisited[i][j] && altVisited[i][j] {
				result = append(result, [][]int{{i, j}}...)
			}
		}
	}
	return result

}

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func bfs(heights [][]int, que *Queue, visited [][]bool) {
	n := len(heights)
	m := len(heights[0])
	for que.size > 0 {
		cur := que.Dequeue()
		for i := range dirs {
			x := cur[0] + dirs[i][0]
			y := cur[1] + dirs[i][1]
			if x < 0 || x >= n || y < 0 || y >= m || visited[x][y] || heights[x][y] < heights[cur[0]][cur[1]] {
				continue
			}
			visited[x][y] = true
			que.Enqueue([]int{x, y})
		}
	}
}

func initVisited(visited [][]bool) {
	for i := range visited {
		visited[i] = make([]bool, len(visited))
	}
}

type Queue struct {
	list [][]int
	size int
}

func (q *Queue) Enqueue(item []int) {
	q.list = append(q.list, item)
	q.size++
}
func (q *Queue) Dequeue() []int {
	item := q.list[0]
	q.list = q.list[1:]
	q.size--
	return item
}

func RunPacAtlantic() {
	heights := [][]int{{1, 1}, {1,1}, {1, 1}}
	fmt.Println(pacificAtlantic(heights))
}
