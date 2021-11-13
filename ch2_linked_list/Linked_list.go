// create a linked list: Node -> Node -> ...
// use dlv to understand the structure of linked list

package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

// display the Linked List
func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

// Insert nodes into list (on the head)
func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

// remove
// to be continued

func main() {
	link := List{}
	link.Insert(1)
	link.Insert(2)
	link.Insert(3)
	link.Insert(4)
	link.Insert(5)
	link.Insert(6)

	link.Display()
	fmt.Printf("Head: %v\n", link.head.key)
	// fmt.Printf("Tail: %v\n", link.tail.key)

	link.Insert(7)

	link.Display()
	fmt.Printf("Head: %v\n", link.head.key)
	// fmt.Printf("Tail: %v\n", link.tail.key)
}
