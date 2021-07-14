package lcs

import "testing"

type TestItems struct {
	Text1    string
	Text2    string
	Expected int
}

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []TestItems{
		{"abcde", "ace", 3},
		{"abc", "def", 0},
		{"abc", "abc", 3},
	}
	for _, te := range tests {
		actual := longestCommonSubsequence(te.Text1, te.Text2)
		if actual != te.Expected {
			t.Errorf("expected %d actual %d", te.Expected, actual)
		}
	}
}

type Tests struct {
	Input    []int
	Expected int
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	tests := []Tests{

		{[]int{3, 4, -1, 0, 6, 2, 3}, 4},
	}
	for _, test := range tests {
		actual := longestIncreasingSubsequence(test.Input)
		if test.Expected != len(actual) {
			t.Errorf("expected %d got %d", test.Expected, len(actual))
		}
	}
}
