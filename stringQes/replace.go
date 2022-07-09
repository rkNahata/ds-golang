package stringQes

import "fmt"

//Given a string a?b?d?f, replace all ‘?’ with a character such that it is not same as its adjacent characters.

func replaceChars(input string)string{
	pre :=- 1
	var preChar,nextChar rune
	for i,ch := range input{
		if pre >=0{
			preChar = ch
		}
		if i+1<=len(input)-1{
			nextChar = rune(input[i+1])
		}
		input[i] = string(getChar(preChar, nextChar))
	}
}

func getChar(pre,next rune)rune{
	alphabets := make([]rune,26)
	if pre !='?'{
		alphabets[pre]--
	}
	if next !='?'{
		alphabets[next]--
	}
	for i:= range alphabets{
		if alphabets[i]>=0{
			return alphabets[i]
		}
	}
	return -1
}
