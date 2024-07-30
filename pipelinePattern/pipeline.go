package main

import "fmt"

func main() {
	//create the initial channel with some data
	data := []int{1, 2, 3, 4, 5, 6, 7}
	input := make(chan int, len(data))
	for _, d := range data {
		input <- d
	}
	close(input)
	//first stage of pipeline:double the input
	doubleOutput := make(chan int)
	go func() {
		defer close(doubleOutput)
		for i := range input {
			doubleOutput <- i * 2
		}
	}()
	//second stage of pipeline:square the doubled value
	squareOutput := make(chan int)
	go func() {
		defer close(squareOutput)
		for d := range doubleOutput {
			squareOutput <- d * d
		}
	}()
	//Third stage of pipeline: printing the square value
	for sqr := range squareOutput {
		fmt.Printf("squared value recived%d\n", sqr)
	}
}
