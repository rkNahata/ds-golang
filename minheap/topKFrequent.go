package minheap

func TopKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	heap := New(make([]int, 0))
	for key, _ := range count {
		heap.Insert1(key, count)
		if heap.Size() > k {
			heap.ExtractMin1(count)
		}
	}
	var result []int
	for heap.Size() > 0 {
		result = append(result, heap.ExtractMin1(count))
	}
	return result

}
