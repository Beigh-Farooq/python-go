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
	tail *Node
	len  int
}

func (l *LinkedList) Add(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		n.next = &n
		l.head = &n
		l.len++
	} else {
		n.next = l.tail.next
		l.tail.next = &n
	}
	l.tail = &n
	l.len++
}

func (l *LinkedList) Display() {
	if l.len == 0 {
		fmt.Println("No Nodes in list")
	}
	p := l.head
	for {
		if p.next == l.head {
			fmt.Printf("-> %v ", p.value)
			break
		}
		fmt.Printf("-> %v ", p.value)
		p = p.next
	}
}
func main() {
	fmt.Println("Circular Linked List In golang")
	l := LinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Display()
}
