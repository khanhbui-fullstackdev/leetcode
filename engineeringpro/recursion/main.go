package main

import (
	"fmt"
	"time"
)

func RunFib() {
	n := 2
	fmt.Printf("Fibonacci number fib(%d):%d", n, fibRecursion(n))
	fmt.Println()

	n = 3
	fmt.Printf("Fibonacci number fib(%d):%d", n, fibRecursion(n))
	fmt.Println()

	n = 4
	fmt.Printf("Fibonacci number fib(%d):%d", n, fibRecursion(n))
	fmt.Println()

	start := time.Now()
	n = 45 // it takes 4.5 seconds to complete
	fmt.Printf("Fibonacci number fib(%d):%d", n, fibRecursion(n))
	fmt.Println()
	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)

	// From n is greater or equal to 45, we must use dynamic programming
	start = time.Now()
	n = 100
	fibonacies := make([]int, n+1)
	//set default value is -1 // fibonacies = [-1 -1 -1 -1 -1,-1]
	fibonacies = setDefaultValueForSlice(fibonacies, -1)
	fibonacies[0] = 0
	fibonacies[1] = 1

	// fmt.Printf("Fibonacci number using dynamic programming topdown fib(%d):%d", n, fibDynamicProgrammingTopdown(n, fibonacies))
	// fmt.Println()
	// elapsed = time.Since(start)
	// fmt.Println("Elapsed:", elapsed)

	fmt.Printf("Fibonacci number using dynamic programming bottomup fib(%d):%d", n, fibDynamicProgrammingTopdown(n, fibonacies))
	fmt.Println()
	elapsed = time.Since(start)
	fmt.Println("Elapsed:", elapsed)
}

/*
Using recursion, everytime fun is called it will create f1(n-1) & f2(n-2)
=> Time complexity O(2^n)
=> Space complexity O(n)
*/
func fibRecursion(n int) int {
	if n < 2 {
		return n
	}
	f1 := fibRecursion(n - 1)
	f2 := fibRecursion(n - 2)
	return f1 + f2
}

/*
Using dynamic programming (top-down)
=> Time complexity O(n)
=> Space complexity O(n)
*/
func fibDynamicProgrammingTopdown(n int, fibonacies []int) int {
	if n < 2 {
		return n
	}

	if fibonacies[n] != -1 {
		return fibonacies[n]
	}
	f1 := fibDynamicProgrammingTopdown(n-1, fibonacies)
	f2 := fibDynamicProgrammingTopdown(n-2, fibonacies)
	fibonacies[n] = f1 + f2
	return fibonacies[n]
}

func setDefaultValueForSlice(nums []int, defaultValue int) []int {
	for index := range nums {
		if nums[index] == 0 {
			nums[index] = defaultValue
		}
	}
	return nums
}

func finDynamicProgrammingBottomup(n int, fibonacies []int) int {
	if n < 2 {
		return n
	}

	for index := 2; index <= n; index++ {
		fibonacies[index] = fibonacies[index-1] + fibonacies[index-2]
	}
	return fibonacies[n]
}

func RunPowerOfThree() {
	n := 27
	fmt.Printf("Is power of three:%v", isPowerOfThree(n))
	fmt.Println()

	n = 0
	fmt.Printf("Is power of three:%v", isPowerOfThree(n))
	fmt.Println()

	n = -1
	fmt.Printf("Is power of three:%v", isPowerOfThree(n))
	fmt.Println()
}

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 || n%3 != 0 {
		return false
	}
	return isPowerOfThree(n / 3)
}

func RunPowerOfFour() {
	n := 16
	fmt.Printf("Is power of four:%v", isPowerOfFour(n))
	fmt.Println()

	n = 5
	fmt.Printf("Is power of four:%v", isPowerOfFour(n))
	fmt.Println()

	n = 1
	fmt.Printf("Is power of four:%v", isPowerOfFour(n))
	fmt.Println()
}

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	} else if n == 0 || n%4 != 0 {
		return false
	}
	return isPowerOfFour(n / 4)
}

func RunIsHappy() {
	n := 19
	fmt.Printf("%d is a happy number:%v", n, isHappy(n))
	fmt.Println()

	n = 2
	fmt.Printf("%d is a happy number:%v", n, isHappy(n))
	fmt.Println()
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	happyNumbers := make([]int, 0, n)
	return isHappyNumber(n, happyNumbers)
}

func isHappyNumber(n int, happyNumbers []int) bool {
	if n == 1 {
		return true
	}
	firstNumber := n % 10
	secondNumber := n / 10
	n = firstNumber*firstNumber + secondNumber*secondNumber
	happyNumbers = append(happyNumbers, n)

	return isHappyNumber(n, happyNumbers)
}
