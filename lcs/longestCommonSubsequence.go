package lcs

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1)==0 || len(text2)==0{
		return 0
	}
	table := make([][]int,len(text1)+1)
	for i:=0;i<=len(text1);i++{
		table[i] = make([]int,len(text2)+1)
	}

	for i:=0;i<len(text1);i++{
		for j:=0;j<len(text2);j++{
			if text1[i]!=text2[j]{
				table[i+1][j+1] = max(table[i][j+1],table[i+1][j])
			}else{
				table[i+1][j+1] = table[i][j]+1
			}
		}
	}
	return table[len(text1)][len(text2)]
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func longestIncreasingSubsequence(input []int)[]int{
	table := make([]int,len(input))
	solution := make([]int,len(input))
	var result []int
	for i:=0;i<len(table);i++{
		table[i]= 1
		solution[i] = i
	}
	for i :=0;i<len(input);i++{
		for j:=0;j<i;j++{
			if input[i]>input[j]{
				if table[j]+1 > table[i]{
					table[i] = table[j]+1
					solution[i] = j
				}

			}
		}
	}
	maxIndex :=0
	for i:=0;i<len(solution);i++{
		if solution[i]>solution[maxIndex]{
			maxIndex = i
		}
	}
	curIndex := solution[maxIndex]
	result = append(result,input[maxIndex])
	for curIndex != maxIndex {
		maxIndex = curIndex
		curIndex = solution[maxIndex]
		result = append(result,input[maxIndex])
	}
	return result
}