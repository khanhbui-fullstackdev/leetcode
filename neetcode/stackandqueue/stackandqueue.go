package main

import (
	"fmt"
	"leetcode/neetcode/stackandqueue/models"
)

func RunImplementStackUsingQueue() {
	myStack := models.MyStack{}

	myStack.Push(1)
	myStack.Push(2)
	fmt.Printf("\n Top:%d", myStack.Top())
	fmt.Printf("\n Pop:%d", myStack.Pop())
	fmt.Printf("\n Is stack empty:%v", myStack.Empty())
	fmt.Println()
}
