package models

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func (l *Node) PrintAllNodes() {
	currentNode := l

	for currentNode != nil {
		fmt.Printf("node:%d ramdom:%v->", currentNode.Val, currentNode.Random)
		currentNode = currentNode.Next
	}
	fmt.Printf("nil")
	fmt.Println()
}
