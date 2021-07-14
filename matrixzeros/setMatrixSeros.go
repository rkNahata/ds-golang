package matrixzeros

import "fmt"

/*
Problem Description: Given a matrix, A of size M x N of 0's and 1's. If an element is 0, set its entire row and column to 0.
w/o using extra space.
For Example:

Input:
[ [1, 0, 1],
[1, 1, 1],
[1, 1, 1] ]

Output:
[ [0, 0, 0],
[1, 0, 1],
[1, 0, 1] ]
*/

func setZeros(matrix [][]int, m, n int) {
	var firstCol, firstRow bool
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRow = true
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstCol = true
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if firstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

	if firstRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	fmt.Println(matrix)
}

func Runner() {
	var matrix = [][]int{{1, 9, 0, 7}, {2, 4, 2, 9}, {8, 4, 6, 0}, {1, 2, 3, 4}}
	setZeros(matrix, len(matrix), len(matrix[0]))
}
