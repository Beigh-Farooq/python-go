package main

import "fmt"

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

//Add Element at last of Linked List
func (l *LinkedList) Add(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

//Add Element at Beginning of Linked List
func (l *LinkedList) add_First(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	n.next = ptr
	l.head = &n
	l.len++
}

//Add Element at any position of LinkedList
func (l *LinkedList) add_Any(val int, pos int) {
	n := Node{}
	n.value = val
	ptr := l.head
	for i := 1; i < pos-1; i++ {
		ptr = ptr.next
	}
	n.next = ptr.next
	ptr.next = &n
	l.len++
}

//Remove Element at Beginning of LinkedList
func (l *LinkedList) delete_First() {
	if l.len == 0 {
		fmt.Println("LinkedList is empty")
		return
	}
	ptr := l.head.next
	l.head = ptr
	l.len--

}

//Remove Element at End of LinkedList
func (l *LinkedList) delete_Last() {
	if l.len == 0 {
		fmt.Println("LinkedList is empty")
		return
	}
	ptr := l.head
	for i := 1; i < l.len-1; i++ {
		ptr = ptr.next
	}
	ptr.next = nil
	l.len--
}

//Display Linked List
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No Nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node:", *ptr)
		ptr = ptr.next
	}
}
func main() {
	fmt.Println("singly Linked List In golang")
	l := LinkedList{}
	fmt.Println("Adding Elements at Last in LinkedList")
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Print()

	fmt.Println("Adding Elements at Beginning of LinkedList")
	l.add_First(12)
	l.Print()

	fmt.Println("Adding Elements at any Position of LinkedList")
	l.add_Any(13, 4)
	l.Print()

	fmt.Println("Delete Element at Beginning of LinkedList")
	l.delete_First()
	l.Print()

	fmt.Println("Delete Element at End of LinkedList")
	l.delete_Last()
	l.Print()

}
