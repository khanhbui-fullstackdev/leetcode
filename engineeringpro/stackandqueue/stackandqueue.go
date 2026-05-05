package main

import (
	"fmt"
	"leetcode/engineeringpro/stackandqueue/models"
	"strings"
)

func RunIsValidParentheses() {
	// s := "()[]{}" // 40 41 91 93 123 125

	s1 := "()"
	fmt.Printf("%s is valid = %v", s1, isValid(s1))
	fmt.Println()

	s2 := "()[]{}"
	fmt.Printf("%s is valid = %v", s2, isValid(s2))
	fmt.Println()

	s3 := "(]"
	fmt.Printf("%s is valid = %v", s3, isValid(s3))
	fmt.Println()

	s4 := "([])"
	fmt.Printf("%s is valid = %v", s4, isValid(s4))
	fmt.Println()

	s5 := "([)]"
	fmt.Printf("%s is valid = %v", s5, isValid(s5))
	fmt.Println()

	s6 := "(((((((("
	fmt.Printf("%s is valid = %v", s6, isValid(s6))
	fmt.Println()
}

func isValid(s string) bool {
	sLength := len(s)
	if sLength == 1 {
		return false
	}

	stackInt := &models.Stack{}

	for _, charS := range s {
		if charS == 40 || charS == 91 || charS == 123 {
			stackInt.Push(int(charS))
		}
		switch charS {

		case 41: // ")"
			openingParentheses, err := stackInt.Pop()
			if openingParentheses != 40 || err != nil {
				return false
			}

		case 93: // "]"
			openingParentheses, err := stackInt.Pop()
			if openingParentheses != 91 || err != nil {
				return false
			}

		case 125: // "}"
			openingParentheses, err := stackInt.Pop()
			if openingParentheses != 123 || err != nil {
				return false
			}
		}
	}
	if stackInt.Count > 0 {
		return false
	}
	return true
}

func RunBackSpaceCompare() {
	// // # = 35
	// s1 := "ab#c"
	// t1 := "ad#c"

	// isEqual1 := backspaceCompare(s1, t1)
	// fmt.Printf("%s & %s is %v", s1, t1, isEqual1)
	// fmt.Println()

	// // # = 35
	// s2 := "ab##"
	// t2 := "c#d#"

	// isEqual2 := backspaceCompare(s2, t2)
	// fmt.Printf("%s & %s is %v", s2, t2, isEqual2)
	// fmt.Println()

	// // # = 35
	// s3 := "a#c"
	// t3 := "b"

	// isEqual3 := backspaceCompare(s3, t3)
	// fmt.Printf("%s & %s is %v", s3, t3, isEqual3)
	// fmt.Println()

	s4 := "xywrrmp"
	t4 := "xywrrmu#p"
	isEqual4 := backspaceCompare(s4, t4)
	fmt.Printf("%s & %s is %v", s4, t4, isEqual4)
	fmt.Println()
}

func backspaceCompare(s string, t string) bool {
	stackS := models.Stack{}
	for _, charS := range s {
		if charS == 35 {
			stackS.Pop()
		} else {
			stackS.Push(int(charS))
		}
	}

	stackT := models.Stack{}
	for _, charT := range t {
		if charT == 35 {
			stackT.Pop()
		} else {
			stackT.Push(int(charT))
		}
	}

	if stackS.Count != stackT.Count {
		return false
	}

	for stackS.Count > 0 {
		itemS, _ := stackS.Pop()
		itemT, _ := stackT.Pop()
		if itemS != itemT {
			return false
		}
	}

	return true
}

func RunPing() {
	recentCounter := models.Constructor()

	fmt.Println(recentCounter.Ping(1))
	fmt.Println(recentCounter.Ping(100))
	fmt.Println(recentCounter.Ping(3001))
	fmt.Println(recentCounter.Ping(3002))
	fmt.Println("Queue:", recentCounter.QueueCounter.Elements)
	fmt.Println()

	rc1 := models.Constructor()
	fmt.Println(rc1.Ping(1))
	fmt.Println(rc1.Ping(2))
	fmt.Println(rc1.Ping(3001))
	fmt.Println("Queue:", rc1.QueueCounter.Elements)
	fmt.Println()

	rc2 := models.Constructor()
	fmt.Println(rc2.Ping(100))
	fmt.Println(rc2.Ping(4000))
	fmt.Println(rc2.Ping(7001))
	fmt.Println("Queue:", rc2.QueueCounter.Elements)
	fmt.Println()

	rc3 := models.Constructor()
	fmt.Println(rc3.Ping(1000))
	fmt.Println(rc3.Ping(4000))
	fmt.Println(rc3.Ping(4001))
	fmt.Println("Queue:", rc3.QueueCounter.Elements)
	fmt.Println()

	rc4 := models.Constructor()
	fmt.Println(rc4.Ping(1))
	fmt.Println(rc4.Ping(4000))
	fmt.Println(rc4.Ping(4001))
	fmt.Println(rc4.Ping(4002))
	fmt.Println("Queue:", rc4.QueueCounter.Elements)
	fmt.Println()
}

func RunCountStudents() {
	// students1 := []int{1, 1, 0, 0}
	// sandwiches1 := []int{0, 1, 0, 1}
	// numberofStudentsAreUnableToEat1 := countStudents(students1, sandwiches1)
	// fmt.Printf("Count students:%d", numberofStudentsAreUnableToEat1)
	// fmt.Println()

	students2 := []int{1, 1, 1, 0, 0, 1}
	sandwiches2 := []int{1, 0, 0, 0, 1, 1}
	numberofStudentsAreUnableToEat2 := countStudents(students2, sandwiches2)
	fmt.Printf("Count students:%d", numberofStudentsAreUnableToEat2)
	fmt.Println()
}

func countStudents(students []int, sandwiches []int) int {
	if len(students) != len(sandwiches) {
		return 0
	}

	queueStudents := models.Queue{}
	for _, student := range students {
		queueStudents.Enqueue(student)
	}

	stackSandwiches := models.Stack{}
	for index := len(sandwiches) - 1; index >= 0; index-- {
		stackSandwiches.Push(sandwiches[index])
	}

	retryCount := 0

	for queueStudents.Count > 0 {
		student := queueStudents.Peek()
		sandwich, _ := stackSandwiches.Peek()
		if student == sandwich {
			queueStudents.Dequeue()
			stackSandwiches.Pop()
			retryCount = 0
		} else {
			queueStudent, _ := queueStudents.Dequeue()
			queueStudents.Enqueue(queueStudent)
			retryCount++
			if retryCount > queueStudents.Count {
				break
			}
		}
	}

	return queueStudents.Count
}

func RunRemoveDuplicates() {
	s1 := "abbaca"
	fmt.Println("Remove duplicate:", removeDuplicates(s1))

	s2 := "azxxzy"
	fmt.Println("Remove duplicate:", removeDuplicates(s2))
}

func removeDuplicates(s string) string {
	lenS := len(s)
	if lenS == 1 {
		return s
	}

	stackS := models.Stack{}
	for index := lenS - 1; index >= 0; index-- {
		intS := (int)(s[index])
		if stackS.Count > 0 {
			item, _ := stackS.Peek()
			if item == intS {
				stackS.Pop()
				continue
			}
		}
		stackS.Push(intS)
	}

	var strBuilder strings.Builder
	strBuilder.Grow(stackS.Count)

	for stackS.Count > 0 {
		item, _ := stackS.Pop()
		strBuilder.WriteRune(rune(item))
	}

	return strBuilder.String()
}

func RunTimeRequiredToBuy() {
	tickets1 := []int{2, 3, 2}
	k1 := 2

	takenTime1 := timeRequiredToBuy(tickets1, k1)
	fmt.Println("Buying time:", takenTime1)

	// tickets2 := []int{5, 1, 1, 1}
	// k2 := 0

	// takenTime2 := timeRequiredToBuy(tickets2, k2)
	// fmt.Println("Buying time:", takenTime2)
}

func timeRequiredToBuy(tickets []int, k int) int {
	queueTicketBuyers := models.QueueObj{}
	for index, ticket := range tickets {
		tickerBuyer := models.NewTicketBuyer(index, ticket)
		queueTicketBuyers.EnqueueObj(tickerBuyer)
	}

	for queueTicketBuyers.Count > 0 {
		// val, _ := queueTicketBuyers.DequeueObj()
		// tickerBuyer := val.(models.TicketBuyer)
		// tickerBuyer.UpdateTicket()

		val := queueTicketBuyers.PeekObj()
		tickerBuyer := val.(models.TicketBuyer)
		if tickerBuyer.Tickets > 0 {
			tickerBuyer.Tickets--
			val, _ = queueTicketBuyers.DequeueObj()
			tickerBuyer = val.(models.TicketBuyer)
			queueTicketBuyers.EnqueueObj(tickerBuyer)
		}
	}

	return 0
}
