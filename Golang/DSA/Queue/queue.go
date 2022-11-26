package main

import "fmt"

//Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

//Enqueue adds a value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//Dequeue deletes a value at the begining
func (q *Queue) Dequeue() {

	q.items = q.items[1:]
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
