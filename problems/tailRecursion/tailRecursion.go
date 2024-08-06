package main

import "fmt"

//There’s no Tail-Call optimization happening in the Go compiler (as of 1.13).
//Tail-call optimization is where you are able to avoid allocating a new stack
//frame for a function because the calling function will simply return the value that
//it gets from the called function. The most common use is tail-recursion,
//where a recursive function written to take advantage of tail-call optimization can use constant stack space.
//find factorial

func main() {
	fmt.Println(Factorial(6))
	fmt.Println(FactorialTail(6))
}

//without tail recursion
func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

//tail-recursion implementation
func FactorialTail(n int) int {
	return calTailFact(n-1, n)
}
func calTailFact(n, current int) int {
	if n == 1 {
		return current
	}
	return calTailFact(n-1, n*current)
}

//fibonacci series
//0,1,1,2,3,5,8,13,21...
func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		res := first
		first, second = second, first+second
		return res
	}

}
