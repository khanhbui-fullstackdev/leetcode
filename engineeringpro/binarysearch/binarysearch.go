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
		} else if nums[middleIndex] < target {
			leftIndex = middleIndex + 1
		} else if nums[middleIndex] > target {
			rightIndex = middleIndex - 1
		}
	}

	return -1
}

var _pickedNumber int

func RunGuessNumber() {
	n := 10
	_pickedNumber = 6
	fmt.Printf("Return picked number:%d", guessNumber(n))
	fmt.Println()

	n = 1
	_pickedNumber = 1
	fmt.Printf("Return picked number:%d", guessNumber(n))
	fmt.Println()

	n = 2
	_pickedNumber = 1
	fmt.Printf("Return picked number:%d", guessNumber(n))
	fmt.Println()
}

func guessNumber(n int) int {
	left := 1
	right := n

	for left <= right {
		middle := (left + right) / 2
		guessRes := guess(middle)
		switch guessRes {
		case 0:
			return middle
		case -1:
			right = middle - 1
		case 1:
			left = middle + 1
		}
	}
	return -1
}

func guess(num int) int {
	if num == _pickedNumber {
		return 0
	} else if num > _pickedNumber {
		return -1
	}
	return 1
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

func RunArrangeCoins() {
	n := 5
	fmt.Printf("Coin:%d => number of complete rows:%d", n, arrangeCoins(n))
	fmt.Println()

	n = 8
	fmt.Printf("Coin:%d => number of complete rows:%d", n, arrangeCoins(n))
	fmt.Println()
}

func arrangeCoins(n int) int {
	minRow := 1
	maxRow := n

	for minRow <= maxRow {
		midRow := (minRow + maxRow) / 2
		totalCoins := calculateNeededCoinsFromRow(midRow)
		if totalCoins == n {
			return totalCoins
		} else if totalCoins > n {
			maxRow = midRow - 1
		} else {
			minRow = midRow + 1
		}
	}
	return maxRow
}

func calculateNeededCoinsFromRow(coin int) int {
	return ((1 + coin) * coin) / 2
}

func RunSearchInsert() {
	nums := []int{1, 3, 5, 6}
	target := 5
	insertedPosition := searchInsert(nums, target)
	fmt.Printf("Nums:%v with target:%d => inserted position:%d", nums, target, insertedPosition)
	fmt.Println()

	nums = []int{1, 3, 5, 6}
	target = 2
	insertedPosition = searchInsert(nums, target)
	fmt.Printf("Nums:%v with target:%d => inserted position:%d", nums, target, insertedPosition)
	fmt.Println()

	nums = []int{1, 3, 5, 6}
	target = 7
	insertedPosition = searchInsert(nums, target)
	fmt.Printf("Nums:%v with target:%d => inserted position:%d", nums, target, insertedPosition)
	fmt.Println()
}

func searchInsert(nums []int, target int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2
		middleVal := nums[middleIndex]
		if middleVal == target {
			return middleIndex
		} else if middleVal > target {
			rightIndex = middleIndex - 1
		} else {
			leftIndex = middleIndex + 1
		}
	}
	return leftIndex
}

func RunSearchRange() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Printf("Nums:%v and target:%d -> search range:%v", nums, target, searchRange(nums, target))
	fmt.Println()

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	fmt.Printf("Nums:%v and target:%d -> search range:%v", nums, target, searchRange(nums, target))
	fmt.Println()

	nums = []int{}
	target = 0
	fmt.Printf("Nums:%v and target:%d -> search range:%v", nums, target, searchRange(nums, target))
	fmt.Println()
}

func searchRange(nums []int, target int) []int {
	rangingIndexes := make([]int, 0, 2)
	minimumIndex := findMinimumIndex(nums, target)
	maximumIndex := findMaximumIndex(nums, target)

	fmt.Println("Minimum index:", minimumIndex)
	fmt.Println("Maximum index:", maximumIndex)

	if minimumIndex > maximumIndex {
		// Nếu minimumIndex > maximumIndex ta return -1, -1 lý do ko tìm thấy target trong array
		minimumIndex = -1
		maximumIndex = -1
	}

	rangingIndexes = append(rangingIndexes, minimumIndex, maximumIndex)
	return rangingIndexes
}

func findMinimumIndex(nums []int, target int) int {
	minimumIndex := 0
	maximunIndex := len(nums) - 1

	for minimumIndex <= maximunIndex {
		middleIndex := (minimumIndex + maximunIndex) / 2
		if nums[middleIndex] >= target {
			maximunIndex = middleIndex - 1
		} else {
			minimumIndex = middleIndex + 1
		}
	}
	return minimumIndex
}

func findMaximumIndex(nums []int, target int) int {
	minimumIndex := 0
	maximunIndex := len(nums) - 1

	for minimumIndex <= maximunIndex {
		middleIndex := (minimumIndex + maximunIndex) / 2
		if nums[middleIndex] <= target {
			minimumIndex = middleIndex + 1
		} else {
			maximunIndex = middleIndex - 1
		}
	}
	return maximunIndex
}

func RunSingleNonDuplicate() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Printf("Nums:%v => single element without duplication:%d", nums, singleNonDuplicate(nums))
	fmt.Println()

	nums = []int{3, 3, 7, 7, 10, 11, 11}
	fmt.Printf("Nums:%v => single element without duplication:%d", nums, singleNonDuplicate(nums))
	fmt.Println()
}

func singleNonDuplicate(nums []int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1
	for leftIndex <= rightIndex {
		if leftIndex == rightIndex {
			return nums[leftIndex]
		}

		midleIndex := (leftIndex + rightIndex) / 2
		if midleIndex%2 != 0 {
			midleIndex = midleIndex - 1
		}
		middleVal := nums[midleIndex]
		nextIndex := midleIndex + 1
		nextVal := nums[nextIndex]

		// a valid pair
		if middleVal == nextVal {
			leftIndex = nextIndex + 1
		} else {
			rightIndex = midleIndex
		}
	}

	return -1
}
