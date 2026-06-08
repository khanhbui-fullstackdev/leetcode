package main

import "fmt"

func RunBinarySearch() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	fmt.Printf("Target:%d with index:%d", target, search(nums, target))
}

func search(nums []int, target int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2
		if nums[middleIndex] == target {
			return middleIndex
		}
		if nums[middleIndex] < target {

		}
	}

	return -1
}

func RunIsPerfectSquare() {
	num := 16
	fmt.Printf("%d is perfect square:%v", num, isPerfectSquare(num))
	fmt.Println()

	num = 14
	fmt.Printf("%d is perfect square:%v", num, isPerfectSquare(num))
	fmt.Println()
}

func isPerfectSquare(num int) bool {
	left := 1
	right := num
	for left <= right {
		middleVal := (left + right) / 2
		powerOf2 := middleVal * middleVal
		if powerOf2 == num {
			return true
		} else if powerOf2 < num {
			left = middleVal + 1
		} else if powerOf2 > num {
			right = middleVal - 1
		}
	}
	return false
}
