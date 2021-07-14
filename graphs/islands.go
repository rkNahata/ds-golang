package graphs

var offset = []int{-1, 0, 1}

func neighbourExists(matrix [][]int, x, y int) bool {
	if x < len(matrix) && x >= 0 && y < len(matrix[0]) && y >= 0 {
		if matrix[x][y] == 1 {
			return true
		}
	}
	return false
}

func dfs(matrix [][]int, x, y int, visited [][]bool) {
	if visited[x][y] {
		return
	}
	visited[x][y] = true
	for _,i := range offset {
		for _,j := range offset {
			if i == 0 && j == 0 {
				continue
			}
			if neighbourExists(matrix, x+i, y+j) {
				dfs(matrix, x+i, y+j, visited)
			}
		}
	}
}

func FindNumberOfClusters(matrix [][]int) int {
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}
	var count int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 && !visited[i][j] {
				count++
				dfs(matrix, i, j, visited)
			}
		}
	}
	return count
}
