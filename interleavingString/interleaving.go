package interleavingString

import "fmt"

func IsInterleave(s1 string, s2 string, s3 string) bool {
	visited := make([][]bool,len(s1)+1)
	for i := range visited {
		visited[i] = make([]bool,len(s2)+1)
	}
	var stack [][]int
	stack = append(stack,[]int{0,0})
	for len(stack)!=0{
		index := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[index[0]][index[1]] = true
		if index[0]==len(s1)&& index[1]==len(s2){
			return true
		}
		if index[0]<len(s1) && !visited[index[0]+1][index[1]] && string(s1[index[0]])==string(s3[index[0]+index[1]]){
			stack = append(stack,[]int{index[0]+1,index[1]})
		}
		if index[1]<len(s2) && !visited[index[0]][index[1]+1] && string(s2[index[1]])==string(s3[index[0]+index[1]]){
			stack = append(stack,[]int{index[0],index[1]+1})
		}
	}
	return false

}

func Runner(){
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	fmt.Print(IsInterleave(s1,s2,s3))
}
