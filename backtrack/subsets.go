package backtrack

import "fmt"

type BK struct {
	result [][]int
	cur    []int
}

func subsets(nums []int) [][]int {
	var bck BK
	bck.cur = []int{}
	bck.result = [][]int{}
	bck.backtrack(nums, 0)
	return bck.result
}

func (bck *BK) backtrack(nums []int, start int) {

	bck.result = append(bck.result, bck.cur)
	for i := start; i < len(nums); i++ {
		bck.cur = append(bck.cur, nums[i])
		bck.backtrack(nums, i+1)
		bck.cur = bck.cur[:len(bck.cur)-1]
	}
}

type Tests struct {
	Input    []int
	Expected [][]int
}

func Runner() {
	tests := Tests{Input: []int{1, 2, 3}}
	fmt.Println(subsets(tests.Input))

}
