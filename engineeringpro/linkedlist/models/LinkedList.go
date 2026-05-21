package models

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	First  *Node // This is the head node
	Last   *Node // We can call it tail node
	Length int
}

// Adds a new node containing the specified value at the start of the Linked List.
// For example: head [1,3,1,8] -> AddFirst(20) -> head [20,1,3,1,8]
func (l *LinkedList) AddFirst(nodeVal int) (*Node, error) {
	if l == nil {
		return nil, errors.New("Linked list cannot be nil")
	}
	if l.First == nil {
		newNode := &Node{Val: nodeVal, Next: nil}
		l.First = newNode
		l.Last = newNode
		l.Length++
		return newNode, nil
	}

	currentHead := l.First
	newNode := &Node{Val: nodeVal, Next: currentHead}
	currentHead.Prev = newNode
	l.First = newNode

	l.Length++
	return newNode, nil
}

// Adds a new node containing the specified value at the end of the Linked List.
// For example: head [1,3,1,8] -> AddLast(30) -> head [1,3,1,8,30]
func (l *LinkedList) AddLast(nodeVal int) (*Node, error) {
	if l == nil {
		return nil, errors.New("Linked list cannot be nil")
	}
	if l.Last == nil {
		newNode := &Node{Val: nodeVal}
		l.Last = newNode
		l.First = newNode
		return newNode, nil
	}

	currentTail := l.Last
	newNode := &Node{Val: nodeVal, Next: nil}
	currentTail.Next = newNode
	newNode.Prev = currentTail
	l.Last = newNode
	return newNode, nil
}

func (l *LinkedList) PrintAllNodes() {
	head := l.First
	fmt.Printf("nil")
	for head != nil {
		fmt.Printf("<->%d", head.Val)
		head = head.Next
	}
	fmt.Printf("<->nil")
	fmt.Println()
}
