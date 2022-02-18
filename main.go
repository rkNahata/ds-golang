package main

import "github.com/rknahata/ds-golang/backtrack"

func main() {

	//trips := [][]int{{2,4,6},{3,2,7},{10,7,9},{8,2,5}}
	//
	//
	//fmt.Println(carPooling(trips,14))


	backtrack.Runner()
}


func carPooling(trips [][]int, capacity int) bool {

	times :=make(map[int]int)
	for _,trip := range trips{
		times[trip[1]] += trip[0]
		times[trip[2]]-=trip[0]
	}
	for _,value := range times{
		capacity -=value
	}
	return capacity>=0

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

