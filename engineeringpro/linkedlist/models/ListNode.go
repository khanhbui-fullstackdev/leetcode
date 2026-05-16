package models

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) PrintAllListNode() {
	currentNode := l
	for currentNode != nil {
		fmt.Printf(" %d ->", currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Println(" nil")
}
