package main

import "fmt"

//queue represents a queue that holds a slice
type Queue struct {
	items []int
}

//Enqueue will add entry at the end
func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

//Dequeue removes a value from the front
//and returns the removed value
func (q *Queue) Dequeue() int {
	if len(q.items) > 0 {
		dequeued := q.items[0]
		q.items = q.items[1:]
		return dequeued
	} else {
		fmt.Println("queue is empty")
	}
	return 0
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(12)
	myQueue.Enqueue(43)
	myQueue.Enqueue(23)
	fmt.Println(myQueue)
	fmt.Println("dequeued :", myQueue.Dequeue())
	fmt.Println("dequeued :", myQueue.Dequeue())
	fmt.Println(myQueue)
}
