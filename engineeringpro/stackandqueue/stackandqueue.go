package main

import (
	"fmt"
	"leetcode/engineeringpro/stackandqueue/models"
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

	stackInt := &models.StackInt{}

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
	if stackInt.Length > 0 {
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
	stackS := models.StackInt{}
	for _, charS := range s {
		if charS == 35 {
			stackS.Pop()
		} else {
			stackS.Push(int(charS))
		}
	}

	stackT := models.StackInt{}
	for _, charT := range t {
		if charT == 35 {
			stackT.Pop()
		} else {
			stackT.Push(int(charT))
		}
	}

	if stackS.Length != stackT.Length {
		return false
	}

	for stackS.Length > 0 {
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
