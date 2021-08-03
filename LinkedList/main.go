package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
	length int
}

func (l *LinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main() {
	myList := LinkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 48}
	node3 := &node{data: 48}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

	fmt.Println(myList.length)
}