package main

import "fmt"

//stack represents stack that holds a slice
type Stack struct {
	items []int
}

//push will add entry at the end
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

//pop will remove a value at the end
//and returns the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	removed := s.items[l]
	s.items = s.items[:l]
	return removed
}
func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(12)
	myStack.Push(13)
	myStack.Push(44)
	myStack.Push(233)
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
}
