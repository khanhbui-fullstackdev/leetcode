package models

import (
	"errors"
	"fmt"
)

type Node struct {
	Id   string
	Val  int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	First   *Node // This is the head node
	Last    *Node // We can call it tail node
	Length  int
	NodeMap map[string]*Node
}

// Adds a new node containing the specified value at the start of the Linked List.
// It works similarly to preappend new element int array/slices
// For example: head [1,3,1,8] -> AddFirst(20) -> head [20,1,3,1,8]
func (l *LinkedList) AddFirst(nodeVal int) (*Node, error) {
	if l == nil {
		return nil, errors.New("Linked list cannot be nil")
	}
	if l.First == nil {
		newNode := &Node{Id: "node#1", Val: nodeVal, Next: nil}
		l.First = newNode
		l.Last = newNode
		l.Length++

		if l.NodeMap == nil {
			l.NodeMap = make(map[string]*Node)
		}

		l.NodeMap["node#1"] = newNode
		return newNode, nil
	}
	newNodeId := fmt.Sprintf("node#%d", l.Length+1)
	newNode := &Node{Id: newNodeId, Val: nodeVal, Next: l.First}
	l.First.Prev = newNode
	l.First = newNode
	l.Length++
	l.NodeMap[newNodeId] = newNode
	return newNode, nil
}

// Adds a new node containing the specified value at the end of the Linked List.
// It works similarly to append new element int array/slices
// For example: head [1,3,1,8] -> AddLast(30) -> head [1,3,1,8,30]
func (l *LinkedList) AddLast(nodeVal int) (*Node, error) {
	if l == nil {
		return nil, errors.New("Linked list cannot be nil")
	}
	if l.Last == nil {
		newNode := &Node{Id: "node#1", Val: nodeVal}
		l.Last = newNode
		l.First = newNode
		l.Length++

		if l.NodeMap == nil {
			l.NodeMap = make(map[string]*Node)
		}
		l.NodeMap["node#1"] = newNode

		return newNode, nil
	}
	newNodeId := fmt.Sprintf("node#%d", l.Length+1)
	newNode := &Node{Id: newNodeId, Val: nodeVal, Next: nil}
	l.Last.Next = newNode
	newNode.Prev = l.Last
	l.Last = newNode
	l.Length++
	l.NodeMap[newNodeId] = newNode
	return newNode, nil
}

func (l *LinkedList) AddNewNodeAfterExistingNode() {

}

func (l *LinkedList) PrintAllNodes() {
	head := l.First
	fmt.Printf("nil")
	for head != nil {
		fmt.Printf("<->%d", head.Val)
		head = head.Next
	}
	fmt.Printf("<->nil")
	fmt.Printf(" || Length:%d", l.Length)
	fmt.Println()
}

func (l *LinkedList) PrintAllNodeMaps() {
	fmt.Printf("\n Print all note maps \n")
	fmt.Println(l.NodeMap)
}

func (l *LinkedList) FindNodeByValue(nodeVal int) (*Node, error) {
	currentNode := l.First
	for currentNode != nil {
		if currentNode.Val == nodeVal {
			return currentNode, nil
		}
		currentNode = currentNode.Next
	}

	return nil, errors.New("Not found")
}

func (l *LinkedList) RemoveNodeByVal(nodeVal int) {
	foundNode, err := l.FindNodeByValue(nodeVal)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// found node is tail
	switch foundNode.Val {
	case l.Last.Val:
		l.RemoveLast()
		return
	case l.First.Val:
		l.RemoveFirst()
		return
	}

	nextNode := foundNode.Next
	previousNode := foundNode.Prev

	previousNode.Next = nextNode
	nextNode.Prev = previousNode
	foundNode = nil
	l.Length--
}

func (l *LinkedList) RemoveFirst() {
	if l == nil || l.Length == 0 {
		fmt.Println("Linked list cannot be null or empty")
		return
	}

	newHead := l.First.Next
	newHead.Prev = nil
	l.First = newHead
	l.Length--
}

func (l *LinkedList) RemoveLast() {
	if l == nil || l.Length == 0 {
		fmt.Println("Linked list cannot be null or empty")
		return
	}

	newTail := l.Last.Prev
	newTail.Next = nil
	l.Last = newTail
	l.Length--
}
