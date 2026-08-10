package main

import (
	"bufio"
	"fmt"
	"leetcode/neetcode/stackandqueue/models"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		showMenu()
		choice := readInput()

		switch choice {
		case 1:
			RunImplementStackUsingQueues()

		case 2:
			RunImplemenQueueUsingStacks()

		case 3:

		case 4:

		case 5:

		case 6:
			RunImplementMinStack()

		case 7:

		case 8:

		case 9:

		case 0:
			fmt.Println("Exit...")

			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func showMenu() {
	fmt.Println("===== 3.Stack and Queue =====")
	fmt.Println("0. Exit")
	fmt.Println("1. 225. Implement Stack using Queues")
	fmt.Println("2. 232. Implement Queue using Stacks")
	fmt.Println("3. 20. Valid Parentheses")
	fmt.Println("4. 844. Backspace String Compare")
	fmt.Println("5. 933. Number of Recent Calls")
	fmt.Println("6. 155. Min Stack")
	fmt.Println("7. 150. Evaluate Reverse Polish Notation")
	fmt.Println("8. ")
	fmt.Println("9. ")
	fmt.Print("Enter your choice: ")
}

func readInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return -1 // invalid
	}

	return num
}

func runQueueTest() {
	queue := &models.QueueInt{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Printf("Queue elements:%v", queue.Elements)
	item, err := queue.Peek()
	if err != nil {
		fmt.Printf("\n Err:%s", err.Error())
	} else {
		fmt.Printf("\n Dequeue:%d", item)
	}

	item, err = queue.Dequeue()
	if err != nil {
		fmt.Printf("\n Err:%s", err.Error())
	} else {
		fmt.Printf("\n Dequeue:%d", item)
	}

	fmt.Printf("\n Queue elements:%v", queue.Elements)
	fmt.Printf("\n Is queue empty:%v", queue.IsEmpty())

	item, err = queue.Dequeue()
	if err != nil {
		fmt.Printf("\n Err:%s", err.Error())
	} else {
		fmt.Printf("\n Dequeue:%d", item)
	}

	item, err = queue.Dequeue()
	if err != nil {
		fmt.Printf("\n Err:%s", err.Error())
	} else {
		fmt.Printf("\n Dequeue:%d", item)
	}
	fmt.Printf("\n Is queue empty:%v", queue.IsEmpty())
	fmt.Println()
}

func runStackTest() {
	stack := &models.StackInt{}
	stack.Push(5)
	stack.Push(8)
	stack.Push(3)

	fmt.Printf("Stack elements:%v", stack.Elements)
	item, err := stack.Pop()
	if err != nil {
		fmt.Printf("\n Err:%s", err.Error())
	} else {
		fmt.Printf("\n Pop:%d", item)
	}

	item, err = stack.Pop()
	if err != nil {
		fmt.Printf("\n Err:%s", err.Error())
	} else {
		fmt.Printf("\n Pop:%d", item)
	}

	fmt.Printf("\n Stack elements:%v", stack.Elements)
	fmt.Printf("\n Is queue empty:%v", stack.IsEmpty())

	item, err = stack.Pop()
	if err != nil {
		fmt.Printf("\n Err:%s", err.Error())
	} else {
		fmt.Printf("\n Pop:%d", item)
	}

	item, err = stack.Pop()
	if err != nil {
		fmt.Printf("\n Err:%s", err.Error())
	} else {
		fmt.Printf("\n Pop:%d", item)
	}
	fmt.Printf("\n Is stack empty:%v", stack.IsEmpty())
	fmt.Println()
}
