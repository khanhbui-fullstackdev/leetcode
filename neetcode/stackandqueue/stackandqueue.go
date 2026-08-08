package main

import (
	"fmt"
	"leetcode/neetcode/stackandqueue/models"
)

func RunImplementStackUsingQueues() {
	myStack := models.MyStack{}

	myStack.Push(1)
	myStack.Push(2)
	fmt.Printf("\n Top:%d", myStack.Top())
	fmt.Printf("\n Pop:%d", myStack.Pop())
	fmt.Printf("\n Is stack empty:%v", myStack.Empty())
	fmt.Println()
}

func RunImplemenQueueUsingStacks() {
	myQueue := models.MyQueue{}
	myQueue.Push(1)
	myQueue.Push(2)

	fmt.Printf("\n Peek:%d", myQueue.Peek())
	fmt.Printf("\n Pop:%d", myQueue.Pop())
	fmt.Printf("\n Is queue empty:%v", myQueue.Empty())
	fmt.Println()

	myQueue = models.MyQueue{}
	myQueue.Push(1)
	myQueue.Push(2)
	myQueue.Push(4)
	myQueue.Push(5)

	fmt.Printf("\n Pop:%d", myQueue.Pop())
	fmt.Printf("\n Pop:%d", myQueue.Pop())
	fmt.Printf("\n Pop:%d", myQueue.Pop())
	fmt.Printf("\n Pop:%d", myQueue.Pop())

	myQueue.Push(7)
	fmt.Printf("\n Pop:%d", myQueue.Pop())
	fmt.Println()
}
