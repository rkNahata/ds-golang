package misc

import (
	"fmt"
	"strconv"
)

func decryptString(text string,freq int){
	for i :=1 ;i<len(text);i+=2 {
		num,_ := strconv.Atoi(string(text[i]))
		if freq<=num{
			fmt.Println(string(text[i-1]))
			return
		}
		freq -= num
	}
	fmt.Println("-1")
}

func Runner(){
	text := "a1b1c3"
	decryptString(text,5)
}

