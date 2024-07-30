package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l linkedlist) printData() {
	toBePrinted := l.head
	for l.length > 0 {
		fmt.Printf("%d", toBePrinted.data)
		toBePrinted = toBePrinted.next
		l.length--
	}
	fmt.Printf("\n")
}
func (l *linkedlist) deleteWithValue(val int) {
	if l.length == 0 {
		return
	}
	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != val {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
func main() {
	mylist := linkedlist{}
	node1 := &node{data: 12}
	node2 := &node{data: 23}
	node3 := &node{data: 65}
	node4 := &node{data: 22}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.printData()
	mylist.deleteWithValue(9822)
	mylist.printData()
}
