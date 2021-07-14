package dfs

import "testing"

type Tests struct {
	PreRequisite [][]int
}

func TestCanFinish(t *testing.T) {

	tests := []Tests{{[][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}}}

	for _, test := range tests {
		Traverse(test.PreRequisite)
	}

}
