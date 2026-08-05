package models

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) PrintAllListNodes() {
	currentNode := l

	for currentNode != nil {
		fmt.Printf("node:%d->", currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Printf("nil")
	fmt.Println()
}
