package hoverhq

/* Solution - 1:
countB(N) = countS(N-1) => there can be only one way to put a building after an empty plot
countS(N) = countS(N-1)+countB(N-1) => there can be 2 ways after an empty plot, either a building or an empty plot
so total ways on one side = countB(N)+countS(N)
*/

func streetPlan(input int) int {
	if input == 1 {
		return 4
	}
	countB, countS := 1, 1
	for i := 2; i <= input; i++ {
		prevB, prevS := countB, countS
		countS, countB = prevB+prevS, prevS
	}
	result := countB + countS
	return result * result
}

/* Solution -2 :
Using fibonacci series. If we carefully examine the output values for one side of the road we get below output
2-3 = fib(4) => fib(input+2)
3-5 = fib(3)
4-8 = fib(4)
5-13 = fib(5)
*/
func streetPlanFib(input int) int {
	var result int
	x, y := 0, 1
	for i := 0; i <= input+2; i++ {
		result = x
		x, y = y, x+y
	}
	return result * result
}

func streetPlanFibRec(input int) int {
	result := fibRec(input)
	return result * result
}

func fibRec(input int) int {
	if input <= 1 {
		return input
	}
	return fibRec(input-1) + fibRec(input-2)
}

func fibGoRoutine(input int)int{
	c := make(chan int)
	var result int
	go fibG(input,c)
	for i:=0;i<input;i++{
		result = <-c
	}
	return result
}

func fibG(input int, c chan int){
	x,y := 0,1
	for i:=0;i<input;i++{
		c <-x
		x,y = y,x+y
	}
	close(c)
}
