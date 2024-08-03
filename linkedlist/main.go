package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

//Insert after Last Node
//insertAtBack(data int) method inserts a new node with the given data at the end of the linked list.

func (l *linkedList) insertAtBack(data int) {
	if l.head == nil {
		l.head = &node{data: data, next: nil}
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &node{data: data, next: nil}
}

//Insert before First Node
//insertAtFront(data int) method inserts a new node with the given data at the beginning (front) of the linked list.
func (l *linkedList) insertAtFront(data int) {
	if l.head == nil {
		l.head = &node{data: data, next: nil}
		return
	}
	next := l.head
	l.head = &node{data: data, next: next}
}

//Insert after Node Value
//insertAfterValue(afterValue, data int) method inserts a new node with the given data immediately after the first occurrence of a specific afterValue in the linked list.
func (l *linkedList) insertAfterValue(afterValue, data int) {
	current := l.head

	for current.next != nil {
		if current.data == afterValue {
			temp := current.next
			current.next = &node{data: data, next: temp}
			return
		}
		current = current.next
	}
	fmt.Printf("can't insert new node after value %d doesn't exists", afterValue)
	fmt.Println()
}

//Insert before Node Value
//insertBeforeValue(beforeValue, data int) method inserts a new node with the given data immediately before the first occurrence of a specific beforeValue in the linked list.

func (l *linkedList) insertBeforeValue(beforeValue, data int) {
	if l.head == nil {
		return
	}
	if l.head.data == beforeValue {
		temp := l.head
		l.head = &node{data: data, next: temp}
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.data == beforeValue {
			temp := current.next
			current.next = &node{data: data, next: temp}
			return
		}
		current = current.next
	}
	fmt.Printf("Cannot insert node, before value %d doesn't exist", beforeValue)
	fmt.Println()
}

//Delete the First Node
//deleteFront() method deletes the first node (the head) of the linked list.

func (l *linkedList) deleteFront() {
	if l.head != nil {
		l.head = l.head.next
		fmt.Printf("head of linked list has been deleted")
		return
	}
}

//Delete the Last Node
//deleteLast() method deletes the last node in the linked list.

func (l *linkedList) deleteLast() {
	if l.head == nil {
		fmt.Printf("linked list is empty \n")
	}
	if l.head.next == nil {
		l.head = nil
		fmt.Printf("last node of linked list has been deleted\n")
		return
	}
	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
	fmt.Printf("last node of linked list has been deleted\n")
}

//Delete after Node Value
//deleteAfterValue(afterValue int) method deletes the node following a node with a specific value (afterValue).

func (l *linkedList) deleteAfterValue(afterValue int) {
	current := l.head

	for current != nil && current.data != afterValue {
		current = current.next
	}
	if current == nil {
		fmt.Println("Node with after value %d doesn't exist\n", afterValue)
		return
	}
	if current.next == nil {
		fmt.Printf("Node with after value %d is the last node\n", afterValue)
		return
	}
	temp := current.next
	fmt.Printf("Node after value %d has data %d and will be deleted", afterValue, temp.data)
	current.next = current.next.next
}

//Delete before Node Value
//deleteBeforeValue(beforeValue int) method deletes the node that precedes a node with a specific value (beforeValue).
func (l *linkedList) deleteBeforeValue(beforeValue int) {
	if l.head == nil || l.head.next == nil {
		fmt.Println("Node with before value %d doesn't exist\n", beforeValue)
		return
	}
	var previous *node
	current := l.head
	for current.next != nil {
		if current.next.data == beforeValue {
			if previous == nil {
				fmt.Printf("Node before value %d has data %d and will be deleted\n", beforeValue, current.data)
				l.head = current.next
			} else {
				fmt.Printf("Node before value %d has data %d and will be deleted\n", beforeValue, current.data)
				previous.next = current.next
			}
			return
		}
		previous = current
		current = current.next
	}
	fmt.Printf("Node before value %d not found\n", beforeValue)
}

//Count Total Nodes
func (l linkedList) Count() (count int) {
	current := l.head
	for current != nil {
		current = current.next
		count++
	}
	return count
}

//Find Node at given Index
//findNodeAt(index int) *Node method searches for and returns the node at a specified index in the linked list.

func (l linkedList) findNodeAt(index int) *node {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	if index <= 0 || index > count {
		return nil
	}
	current = l.head
	for count = 1; count < index; count++ {
		current = current.next
	}
	return current
}

//Print the Linked List
func (l linkedList) print() {
	for l.head != nil {
		fmt.Printf("%d ->", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}

func main() {
	// Create an empty linked list
	myList := linkedList{}

	// Insert some data at the beginning
	myList.insertAtFront(10)
	myList.insertAtFront(20)
	myList.insertAtFront(30)
	myList.insertAfterValue(10, 40)
	// Print the contents of the linked list
	fmt.Println("Linked List Contents:")
	myList.print()

	// Count the nodes in the linked list
	count := myList.Count()
	fmt.Printf("Total number of nodes: %d\n", count)

	// Find and print a node at a specific index
	indexToFind := 4
	foundNode := myList.findNodeAt(indexToFind)
	if foundNode != nil {
		fmt.Printf("Node at index %d has data: %d\n", indexToFind, foundNode.data)
	} else {
		fmt.Printf("Node at index %d not found\n", indexToFind)
	}

}

// type node struct {
// 	data int
// 	next *node
// }

// type linkedlist struct {
// 	head   *node
// 	length int
// }

// func (l *linkedlist) prepend(n *node) {
// 	second := l.head
// 	l.head = n
// 	l.head.next = second
// 	l.length++
// }
// func (l linkedlist) printData() {
// 	toBePrinted := l.head
// 	for l.length > 0 {
// 		fmt.Printf("%d", toBePrinted.data)
// 		toBePrinted = toBePrinted.next
// 		l.length--
// 	}
// 	fmt.Printf("\n")
// }
// func (l *linkedlist) deleteWithValue(val int) {
// 	if l.length == 0 {
// 		return
// 	}
// 	if l.head.data == val {
// 		l.head = l.head.next
// 		l.length--
// 		return
// 	}
// 	previousToDelete := l.head
// 	for previousToDelete.next.data != val {
// 		if previousToDelete.next.next == nil {
// 			return
// 		}
// 		previousToDelete = previousToDelete.next
// 	}
// 	previousToDelete.next = previousToDelete.next.next
// 	l.length--
// }
// func main() {
// 	mylist := linkedlist{}
// 	node1 := &node{data: 12}
// 	node2 := &node{data: 23}
// 	node3 := &node{data: 65}
// 	node4 := &node{data: 22}
// 	mylist.prepend(node1)
// 	mylist.prepend(node2)
// 	mylist.prepend(node3)
// 	mylist.prepend(node4)
// 	mylist.printData()
// 	mylist.deleteWithValue(9822)
// 	mylist.printData()
// }
