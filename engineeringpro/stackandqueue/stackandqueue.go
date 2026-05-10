package main

import (
	"fmt"
	"leetcode/engineeringpro/stackandqueue/models"
	"strconv"
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
	fmt.Println()

	tickets2 := []int{5, 1, 1, 1}
	k2 := 0

	takenTime2 := timeRequiredToBuy(tickets2, k2)
	fmt.Println("Buying time:", takenTime2)
	fmt.Println()
}

func timeRequiredToBuy(tickets []int, k int) int {
	queueTicketBuyers := models.QueueObj{}
	for index, ticket := range tickets {
		tickerBuyer := models.NewTicketBuyer(index, ticket)
		queueTicketBuyers.EnqueueObj(tickerBuyer)
	}

	takenTime := 0
	for queueTicketBuyers.Count > 0 {
		takenTime++

		val := queueTicketBuyers.PeekObj()
		tickerBuyer := val.(models.TicketBuyer)
		tickerBuyer.Tickets--
		queueTicketBuyers.DequeueObj()

		if tickerBuyer.Tickets == 0 {
			if tickerBuyer.Index == k {
				break
			}
		} else {
			queueTicketBuyers.EnqueueObj(tickerBuyer)
		}
	}
	return takenTime
}

func RunCalPoints() {
	operations1 := []string{"5", "2", "C", "D", "+"}
	scores1 := calPoints(operations1)
	fmt.Println("Score:", scores1)

	operations2 := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	scores2 := calPoints(operations2)
	fmt.Println("Score:", scores2)

	operations3 := []string{"1", "C"}
	scores3 := calPoints(operations3)
	fmt.Println("Score:", scores3)
}

func calPoints(operations []string) int {
	score := 0
	stackScore := models.Stack{}

	for _, operation := range operations {
		number, err := strconv.Atoi(operation)
		if err != nil {
			switch operation {
			case "C":
				stackScore.Pop()
			case "D":
				previousScore, _ := stackScore.Peek()
				stackScore.Push(previousScore * 2)
			case "+":
				previousScore1, _ := stackScore.Pop()
				previousScore2, _ := stackScore.Peek()

				stackScore.Push(previousScore1)
				totalScore := previousScore1 + previousScore2
				stackScore.Push(totalScore)
			}
		} else {
			stackScore.Push(number)
		}
	}

	for stackScore.Count > 0 {
		item, _ := stackScore.Pop()
		score += item
	}

	return score
}

func RunMakeGood() {
	// a:97 A:65 b:98 B:66 c:99 C:67
	s1 := "leEeetcode"
	fmt.Printf("s1: %s -> after process = %s", s1, makeGood(s1))
	fmt.Println()

	s2 := "abBAcC"
	fmt.Printf("s2: %s -> after process = %s", s2, makeGood(s2))
	fmt.Println()

	s3 := "abBAcC"
	fmt.Printf("s2: %s -> after process = %s", s3, makeGood(s3))
	fmt.Println()
}

func makeGood(s string) string {
	sLength := len(s)
	if sLength == 1 {
		return s
	}
	stack := models.Stack{}
	for _, charS := range s {
		if stack.Count == 0 {
			stack.Push((int(charS)))
			continue
		}
		item, _ := stack.Peek()
		if absInt(item-int(charS)) == 32 {
			stack.Pop()
		} else {
			stack.Push((int(charS)))
		}
	}
	var stringBuilder strings.Builder
	stringBuilder.Grow(stack.Count)
	nums := make([]int, 0, stack.Count)
	for stack.Count > 0 {
		item, _ := stack.Pop()
		nums = append(nums, item)
	}

	for index := len(nums) - 1; index >= 0; index-- {
		runeItem := rune(nums[index])
		stringBuilder.WriteRune(runeItem)
	}

	return stringBuilder.String()
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func RunMinLength() {
	// s33 := "ABCD"
	// fmt.Printf("A:%d", s33[0])   //A:65
	// fmt.Printf("\nB:%d", s33[1]) // B:66
	// fmt.Printf("\nC:%d", s33[2]) //C:67
	// fmt.Printf("\nD:%d", s33[3]) //D:68

	s1 := "ABFCACDB"
	minLength1 := minLength(s1)
	fmt.Printf("\n Min length:%d", minLength1)
	fmt.Println()

	s2 := "ACBBD"
	minLength2 := minLength(s2)
	fmt.Printf("\n Min length:%d", minLength2)
	fmt.Println()

	s1496 := "VKBAJBOYY"
	minLength1496 := minLength(s1496)
	fmt.Printf("\n Min length:%d", minLength1496)
	fmt.Println()
}

func minLength(s string) int {
	stack := models.Stack{}

	for _, charS := range s {
		if stack.Count == 0 {
			stack.Push(int(charS))
			continue
		}
		item, _ := stack.Peek()
		if (charS == 66 && item == 65) || (item == 67 && charS == 68) {
			stack.Pop()
		} else {
			stack.Push(int(charS))
		}
	}
	fmt.Println("Stack:", stack)

	return stack.Count
}

func RunBasicCalculatorII() {
	// +:43 -:45 *:42 /:47

	s1 := "3+2*2" // 7 ✅
	fmt.Printf("Basic calculator II of %s  = %d", s1, calculateII(s1))
	fmt.Println()

	s1 = "3-2*2" // -1 ✅
	fmt.Printf("Basic calculator II of %s  = %d", s1, calculateII(s1))
	fmt.Println()

	s1 = "-3-2*2+17" // 10 ✅
	fmt.Printf("Basic calculator II of %s  = %d", s1, calculateII(s1))
	fmt.Println()

	s1 = "-3+17" // 10 ✅
	fmt.Printf("Basic calculator II of %s  = %d", s1, calculateII(s1))
	fmt.Println()

	s2 := " 3/2 " // 0✅
	fmt.Printf("Basic calculator II of %s  = %d", s2, calculateII(s2))
	fmt.Println()

	s3 := " 3+50 / 2 " // 28 ✅
	fmt.Printf("Basic calculator II of %s  = %d", s3, calculateII(s3))
	fmt.Println()

	s3 = " 525 / 25 " // 21
	fmt.Printf("Basic calculator II of %s  = %d", s3, calculateII(s3))
	fmt.Println()

	s3 = " 3*50 / 10 "
	fmt.Printf("Basic calculator II of %s  = %d", s3, calculateII(s3))
	fmt.Println()
}

// +:43 -:45 *:42 /:47
func calculateII(s string) int {
	s = strings.TrimSpace(s)
	lenS := len(s)
	if lenS == 1 {
		return int(s[0])
	}
	// loop through string s,
	// step 1 check if charS is empty space 32 -> ' ' -> continue
	// step 2 build number
	// step 3 when charS is a signedOperation '+', '-', '*', '/' then we prioritize * and / then we execute the operation and push the result into stack
	currentSign := 43
	buildNumber := 0
	stack := models.Stack{}
	for index, charS := range s {
		if charS == 32 {
			continue
		}
		if isANumber(charS) {
			digit := int(charS - '0')
			buildNumber = buildNumber*10 + digit
		}
		if !isANumber(charS) || index == lenS-1 {
			switch currentSign {
			case 43:
				stack.Push(buildNumber)
			case 45:
				stack.Push(-buildNumber)
			case 42:
				preNumber, _ := stack.Pop()
				stack.Push(preNumber * buildNumber)
			case 47:
				preNumber, _ := stack.Pop()
				stack.Push(preNumber / buildNumber)
			}
			// set current sign to default '+'
			currentSign = int(charS)
			buildNumber = 0
		}
	}

	fmt.Println("Stack:", stack)

	total := 0
	for stack.Count > 0 {
		item, _ := stack.Pop()
		total += item
	}
	return total
}

func isANumber(charS rune) bool {
	return charS != 43 && charS != 45 && charS != 42 && charS != 47 && charS != 40 && charS != 41
}

func RunCalculate() {
	// +:43 -:45 (:40 ):41

	s1 := "1 + 1" // 2 ✅
	fmt.Printf("Basic calculator of %s  = %d", s1, calculate(s1))
	fmt.Println()

	s1 = " 2-1 + 2 " // 3 ✅
	fmt.Printf("Basic calculator of %s  = %d", s1, calculate(s1))
	fmt.Println()

	s1 = " 2-1 + (2+10) " // 13 ✅
	fmt.Printf("Basic calculator of %s  = %d", s1, calculate(s1))
	fmt.Println()

	s1 = "(1+(4+5+2)-3)+(6+8)" // 23 ✅
	fmt.Printf("Basic calculator of %s  = %d", s1, calculate(s1))
	fmt.Println()
}

func calculate(s string) int {
	s = strings.TrimSpace(s)

	lenS := len(s)
	if lenS == 1 {
		return int(s[0] - '0')
	}
	buildNumber := 0
	total := 0
	currentSign := 1 // '1' is plus sign; '-1' is negative sign
	stack := models.Stack{}
	for index, charS := range s {
		if charS == 32 {
			continue
		}
		if isANumber(charS) {
			digit := int(charS - '0')
			buildNumber = buildNumber*10 + digit
			if index == lenS-1 {
				total += currentSign * buildNumber
			}
		} else {
			switch charS {
			case 43: // '+'
				total += buildNumber
				currentSign = 1

			case 45: // '-'
				total += buildNumber
				currentSign = -1

			case 40: // '('
				// push number (1) and sign (2) into stack
				stack.Push(total)
				stack.Push(currentSign)
				total = 0
				currentSign = 1

			case 41: // ')'
				if currentSign == 1 {
					total += buildNumber
				} else {
					total -= buildNumber
				}
				// first pop => get sign from stack
				item, _ := stack.Pop()
				total *= item

				// second pop => get sign from
				item2, _ := stack.Pop()
				total += item2
			}
			buildNumber = 0
		}
	}

	return total
}
