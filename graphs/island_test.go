package graphs

import "testing"

func TestFindNumberOfClusters(t *testing.T) {

	type Tests struct {
		Matrix   [][]int
		Expected int
	}
	tests := []Tests{{Matrix: [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1}}, Expected: 3,
	}, {Matrix: [][]int{
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 0},
		{0, 0, 0, 0, 1}}, Expected: 4,
	}, {Matrix: [][]int{{1}}, Expected: 1,
	}, {Matrix: [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}, Expected: 0,
	}}
	for i, test := range tests {
		actual := FindNumberOfClusters(test.Matrix)
		if test.Expected != actual {
			t.Errorf("[%v] Expected %v Got %v", i+1, test.Expected, actual)
		}
	}

}