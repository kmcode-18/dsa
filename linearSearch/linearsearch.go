package main

import "fmt"

//return index position of target or -1 if not found
func linearSearch(list []int, target int) int {
	for i, v := range list {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	list := []int{12, 3, 34, 5, 66, 78, 32, 11}
	fmt.Println(linearSearch(list, 664))
}
