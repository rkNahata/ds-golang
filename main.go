package main

import (
	"github.com/rknahata/ds-golang/graphs"
	"sort"
)

func main() {

	//interleaver.Runner()
	graphs.RunPacAtlantic()
}

func carPooling(trips [][]int, capacity int) bool {

	times := make(map[int]int)
	for _, trip := range trips {
		times[trip[1]] += trip[0]
		times[trip[2]] -= trip[0]
	}
	for _, value := range times {
		capacity -= value
	}
	return capacity >= 0

	//sort.SliceStable(trips,func(a,b int)bool{
	//	return trips[a][1]<trips[b][1]
	//})
	//fmt.Println(trips)
	//tracker :=make([]int,2)
	//tracker[0]=trips[0][0]
	//tracker[1]=trips[0][2]
	//capacity -= trips[0][0]
	//for i:=1;i<len(trips);i++{
	//
	//	for j :=i-1;j>=0;j--{
	//		if trips[j][2]<=trips[i][1]{
	//			capacity +=trips[j][0]
	//		}
	//	}
	//
	//	if trips[i][0]>capacity{
	//		return false
	//	}
	//
	//	if tracker[1]>trips[i][2]{
	//		capacity -= trips[i][0]
	//		tracker[0]=trips[i][0]
	//		tracker[1]= trips[i][2]
	//	}
	//
	//
	//
	//}
	//return true
}

func minDeletions(s string) int {
	freq := make([]int, 26)
	var maxFreqCount int
	var deleted int
	for i := range s {
		freq[s[i]-'a']++
		maxFreqCount++
	}
	sort.Ints(freq)
	for i := 25; i >= 0; i-- {
		if freq[i] > maxFreqCount {
			deleted += freq[i] - maxFreqCount
			freq[i] = maxFreqCount
		}
		maxFreqCount = max(0, freq[i]-1)
	}

	return deleted
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
