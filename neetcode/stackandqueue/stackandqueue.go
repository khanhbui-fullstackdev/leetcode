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

func RunImplementMinStack() {
	// stack = [-2,0,-3]
	// minStack = [-2,-3]
	minStack := models.NewMinStackV2()
	minStack.Push(-2)
	fmt.Printf("Min value:%d", minStack.GetMin())
	fmt.Println()

	minStack.Push(0)
	minStack.Push(-3) // minStack = [-2, 0, -3]

	fmt.Printf("Min value:%d", minStack.GetMin())
	fmt.Println()

	minStack.Pop()
	fmt.Printf("Top:%d", minStack.Top())
	fmt.Println()

	fmt.Printf("Min value:%d", minStack.GetMin())
	fmt.Println()

	minStack = models.NewMinStackV2()
	minStack.Push(5)
	fmt.Printf("Min value:%d", minStack.GetMin())
	fmt.Println()

	minStack.Push(3)
	minStack.Push(4) // minStack = [-2, 0, -3]
	minStack.Push(2)

	minStack.PrintAllStackVals()
	minStack.PrintAllMinStackVals()

	fmt.Println("\n *** Test case 30 ***")
	// ["MinStack","push","push","push","getMin","top","pop","getMin"]
	// [[],[-2],[0],[-1],[],[],[],[]]

	minStack = models.NewMinStackV2()
	minStack.Push(-2)
	minStack.PrintAllMinStackVals()
	minStack.PrintAllStackVals()

	minStack.Push(0)
	minStack.PrintAllMinStackVals()
	minStack.PrintAllStackVals()

	minStack.Push(-1)
	minStack.PrintAllMinStackVals()
	minStack.PrintAllStackVals()

	fmt.Printf("Min value:%d", minStack.GetMin())
	fmt.Println()

	fmt.Printf("Top value:%d", minStack.Top())
	fmt.Println()

	fmt.Printf("Get min value:%d", minStack.GetMin())
	fmt.Println()

	fmt.Println("\n *** Test case 45 ***")
	/*
		["MinStack","push","push","push","getMin","pop","getMin","pop","getMin","pop","push","push","push","getMin","pop","top","getMin","pop","getMin","pop"]
		[[],[0],[1],[0],[],[],[],[],[],[],[-2],[-1],[-2],[],[],[],[],[],[],[]]
	*/
	minStack = models.NewMinStackV2()
	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)

	minStack.PrintAllStackVals()
	minStack.PrintAllMinStackVals()

	fmt.Printf("Get min value:%d", minStack.GetMin())
	minStack.Pop()
	minStack.PrintAllStackVals()
	minStack.PrintAllMinStackVals()

	fmt.Printf("Get min value:%d", minStack.GetMin())
	minStack.Pop()
	minStack.PrintAllStackVals()
	minStack.PrintAllMinStackVals()
	fmt.Println()
}
