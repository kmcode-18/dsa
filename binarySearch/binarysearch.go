package main

import (
	"fmt"
)

//c1
//must be sorted list
//return index position of target or -1 if not found

func binarySearch(list []int, target int) int {
	first := 0
	last := len(list) - 1
	for first <= last {
		mid := (first + last) / 2
		if list[mid] == target {
			return mid
		} else if list[mid] > target {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return -1
}

//must be sorted list
//return bool true if target exists and false if don't exists
func binarySearchRecursive(list []int, target int) bool {
	if len(list) == 0 {
		return false
	} else {
		midPoint := len(list) / 2
		if list[midPoint] == target {
			return true
		} else if list[midPoint] > target {
			return binarySearchRecursive(list[:midPoint], target)
		} else {
			return binarySearchRecursive(list[midPoint+1:], target)
		}
	}
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(binarySearchRecursive(list, 0))
}
