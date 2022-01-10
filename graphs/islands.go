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
	for _, i := range offset {
		for _, j := range offset {
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

type GridData struct {
	grid    [][]int
	r, c    int
	count   int
	maxArea int
	land    int
}

func maxAreaOfIsland(grid [][]int) int {
	gridData := GridData{grid: grid, r: len(grid) - 1, c: len(grid[0]) - 1, count: 0, maxArea: 0, land: 1}
	if gridData.r+1==1 && gridData.c+1==1 {
		return grid[0][0]
	}

	for i := 0; i < len(gridData.grid); i++ {
		for j := 0; j < len(gridData.grid[0]); j++ {
			gridData.area(i, j)
			if gridData.count > gridData.maxArea {
				gridData.maxArea = gridData.count
			}
			gridData.count = 0
		}
	}
	return gridData.maxArea
}

func (g *GridData) area(i, j int) {
	if g.grid[i][j] == g.land {
		g.grid[i][j] = -1
		g.count += 1
		if i > 0 {
			g.area(i-1, j)
		}
		if j > 0 {
			g.area(i, j-1)
		}
		if i < g.r {
			g.area(i+1, j)
		}
		if j < g.c {
			g.area(i, j+1)
		}
	}
}
