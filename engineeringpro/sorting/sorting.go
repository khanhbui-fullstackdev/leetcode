package main

import (
	"cmp"
	"fmt"
	"slices"
)

func RunTargetIndicates() {
	nums := []int{1, 2, 5, 2, 3}
	target := 2

	indices := targetIndicesUsingBinarySearch(nums, target)
	fmt.Printf("Nums:%v and target:%d -> indicies:%v", nums, target, indices)
	fmt.Println()

	nums = []int{1, 2, 5, 2, 3}
	target = 3

	indices = targetIndicesUsingBinarySearch(nums, target)
	fmt.Printf("Nums:%v and target:%d -> indicies:%v", nums, target, indices)
	fmt.Println()

	nums = []int{1, 2, 5, 2, 3}
	target = 5

	indices = targetIndicesUsingBinarySearch(nums, target)
	fmt.Printf("Nums:%v and target:%d -> indicies:%v", nums, target, indices)
	fmt.Println()
}

// Linear solution
// Time complexity =  Sorted O(n*logn) + O(n) => total = O(n log n) since nlogn > n
// Space complexity = O(n)
func targetIndices(nums []int, target int) []int {
	indices := make([]int, 0, len((nums)))

	if !slices.IsSorted(nums) {
		slices.SortFunc(nums, func(a, b int) int {
			return cmp.Compare(a, b)
		})
	}

	for index, num := range nums {
		if num == target {
			indices = append(indices, index)
		}
	}

	return indices
}

// Binary search solution
// Time complexity =  Sorted O(n*logn) + findMinimumIndex = O(log(2)n)  + findMaximumIndex = O(log(2)n) => Total = O(n*logn) + 2 * (O(log(2)n)) = O(nlogn)
// Space complexity = O(n)
func targetIndicesUsingBinarySearch(nums []int, target int) []int {
	if !slices.IsSorted(nums) {
		slices.SortFunc(nums, func(a, b int) int {
			return cmp.Compare(a, b)
		})
	}
	minimumIndex := findMinimumIndex(nums, target)
	maximumIndex := findMaximumIndex(nums, target)
	if minimumIndex > maximumIndex {
		return []int{}
	}
	indices := make([]int, 0, maximumIndex)
	for index := minimumIndex; index <= maximumIndex; index++ {
		indices = append(indices, index)
	}
	return indices
}

// Find the minimum index match condition middle[leftIndex] >= target
// nums = [1,2,2,3,5], target = 2
// middleValue[1] = 2 (2 >= 2) => minimumIndex = 2
// Time complexity = O(log2N)
func findMinimumIndex(nums []int, target int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2
		middleValue := nums[middleIndex]
		if middleValue >= target {
			rightIndex = middleIndex - 1
		} else {
			leftIndex = middleIndex + 1
		}
	}
	return leftIndex
}

// Find the maximum index match condition middle[rightIndex] <= target
// nums = [1,2,2,3,5], target = 2
// middleValue[3] = 2 (2 <= 2) => maximumIndex = 3
// Time complexity = O(log2N)
func findMaximumIndex(nums []int, target int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2
		middleValue := nums[middleIndex]
		if middleValue <= target {
			leftIndex = middleIndex + 1
		} else {
			rightIndex = middleIndex - 1
		}
	}
	return rightIndex
}
