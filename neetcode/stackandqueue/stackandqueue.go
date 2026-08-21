package main

import (
	"fmt"
	"leetcode/neetcode/stackandqueue/models"
	"strconv"
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

func RunDailyTemperatures() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	waitingDays := dailyTemperaturesUsingStack(temperatures)
	fmt.Printf("Waiting days:%v", waitingDays)
	fmt.Println()

	temperatures = []int{30, 40, 50, 60}
	waitingDays = dailyTemperaturesUsingStack(temperatures)
	fmt.Printf("Waiting days:%v", waitingDays)
	fmt.Println()

	temperatures = []int{30, 60, 90}
	waitingDays = dailyTemperaturesUsingStack(temperatures)
	fmt.Printf("Waiting days:%v", waitingDays)
	fmt.Println()
}

/*
	always starts with brute solution
	            0   1   2   3   4   5   6   7

temperatures = [73, 74, 75, 71, 69, 72, 76, 73]   len(temperatures) = 8

					            i
						                j
	    output = []
	    for i=0;i<len(temperatures)<;i++{
	        for j=i+1;j<len(temperatures);j++{
	            i = 0
	            {
	                j = 1
	                (73<74){
	                    waitingDays:=j-i=1
	                    output=[1]
	                    i++ // 1
	                    continue
	                }
	            }
	            i = 1
	            {
	               j = 2
	               (74<75){
	                     waitingDays:=j-i=1
	                     output=[1,1]
	                     i++ // 2
	                     continue
	               }
	            }
	            i = 2
	            {
	                j = 4
	                (75<71)
	                j = 5
	                j = 6
	                {
	                     (75<76)
	                     {
	                        waitingDays:=j-i=4
	                        output=[1,1,4]
	                        i++//3
	                        continue
	                     }
	                }
	            }
	            i = 3{

	            }
	        }
	    }
*/
func dailyTemperaturesBruteForce(temperatures []int) []int {
	lenTemperatures := len(temperatures)
	numberofWaitingDays := make([]int, 0, lenTemperatures)

	for index := 0; index < lenTemperatures; index++ {
		if index == lenTemperatures-1 {
			numberofWaitingDays = append(numberofWaitingDays, 0)
		}
		for jIndex := index + 1; jIndex < lenTemperatures; jIndex++ {
			if temperatures[index] < temperatures[jIndex] {
				numberofWaitingDays = append(numberofWaitingDays, jIndex-index)
				break
			} else if jIndex == lenTemperatures-1 {
				numberofWaitingDays = append(numberofWaitingDays, 0)
			}
		}
	}

	return numberofWaitingDays
}

/*
 0   1   2   3   4   5   6   7   len = 8
[73, 74, 75, 71, 69, 72, 76, 73]  input
[1,  1,  4,  2,  1,  1,  0,  0]   output

main idea: compare 2 temperature and store index the warmer one

*/

func dailyTemperaturesUsingStack(temperatures []int) []int {
	lenTemperatures := len(temperatures)
	numberofWaitingDays := make([]int, lenTemperatures)

	stack := models.StackInt{}
	for index, temperature := range temperatures {
		for stack.Size > 0 {
			previousIndex, _ := stack.Peek()
			if temperatures[previousIndex] < temperature {
				previousIndex, _ = stack.Pop()
				numberofWaitingDays[previousIndex] = index - previousIndex
			} else {
				stack.Push(index)
				break
			}
		}

		if stack.IsEmpty() {
			stack.Push(index)
		}
	}

	if stack.Size > 0 {
		clear(stack.Elements)
		stack.Elements = nil
		stack.Size = 0
	}

	return numberofWaitingDays
}

func RunEvalRPN() {
	tokens := []string{"2", "1", "+", "3", "*"}
	valueOfExpression := evalRPN(tokens)
	fmt.Printf("Token:%v => Value:%d", tokens, valueOfExpression)
	fmt.Println()

	tokens = []string{"4", "13", "5", "/", "+"}
	valueOfExpression = evalRPN(tokens)
	fmt.Printf("Token:%v => Value:%d", tokens, valueOfExpression)
	fmt.Println()

	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	valueOfExpression = evalRPN(tokens)
	fmt.Printf("Token:%v => Value:%d", tokens, valueOfExpression)
	fmt.Println()
}

func evalRPN(tokens []string) int {
	valueOfExpression := 0

	stack := models.StackInt{}

	for _, token := range tokens {
		if isValidNumber(token) {
			num, err := strconv.Atoi(token)
			if err != nil {
				continue
			}
			stack.Push(num)
		} else {
			switch token {
			case "+":
				item1, _ := stack.Pop()
				item2, _ := stack.Pop()
				stack.Push(item2 + item1)

			case "-":
				item1, _ := stack.Pop()
				item2, _ := stack.Pop()
				stack.Push(item2 - item1)

			case "*":
				item1, _ := stack.Pop()
				item2, _ := stack.Pop()
				stack.Push(item2 * item1)

			case "/":
				item1, _ := stack.Pop()
				item2, _ := stack.Pop()
				stack.Push(item2 / item1)
			}
		}
	}

	for stack.Size > 0 {
		item, _ := stack.Pop()
		valueOfExpression += item
	}

	return valueOfExpression
}

func isValidNumber(s string) bool {
	return s != "+" && s != "-" && s != "*" && s != "/"
}
