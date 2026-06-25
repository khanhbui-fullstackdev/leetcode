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

func RunSmallerNumbersThanCurrent() {
	nums := []int{8, 1, 2, 2, 3}
	countSmallerNumbers := smallerNumbersThanCurrent(nums)
	fmt.Printf("Count smaller numbers:%v", countSmallerNumbers)
	fmt.Println()

	nums = []int{6, 5, 4, 8}
	countSmallerNumbers = smallerNumbersThanCurrent(nums)
	fmt.Printf("Count smaller numbers:%v", countSmallerNumbers)
	fmt.Println()

	nums = []int{7, 7, 7, 7}
	countSmallerNumbers = smallerNumbersThanCurrent(nums)
	fmt.Printf("Count smaller numbers:%v", countSmallerNumbers)
	fmt.Println()
}

func smallerNumbersThanCurrent(nums []int) []int {
	lenNums := len(nums)
	// Copy all data from slice nums -> slices sortedNums
	sortedNums := append([]int(nil), nums...)

	if !slices.IsSorted(sortedNums) {
		slices.SortFunc(sortedNums, func(a, b int) int {
			return cmp.Compare(a, b)
		})
	}
	sortedMapNums := make(map[int]int, lenNums)

	for index, num := range sortedNums {
		_, existed := sortedMapNums[num]
		if !existed {
			sortedMapNums[num] = index
		}
	}

	for index, num := range nums {
		nums[index] = sortedMapNums[num]
	}

	return nums
}

func RunRelativeSortArray() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	sortedArray := relativeSortArrayUsingArr(arr1, arr2)
	fmt.Printf("Relative sort array:%v", sortedArray)
	fmt.Println()

	arr1 = []int{28, 6, 22, 8, 44, 17}
	arr2 = []int{22, 28, 8, 6}
	sortedArray = relativeSortArrayUsingArr(arr1, arr2)
	fmt.Printf("Relative sort array:%v", sortedArray)
	fmt.Println()
}

func relativeSortArrayUsingMap(arr1 []int, arr2 []int) []int {
	lenArr1 := len(arr1)

	outputArr := make([]int, 0, lenArr1)
	numFrequencies := make(map[int]int, lenArr1)

	for _, num := range arr1 {
		numFrequencies[num]++
	}

	for _, item := range arr2 {
		frequencies, existed := numFrequencies[item]
		if existed {
			for index := 1; index <= frequencies; index++ {
				outputArr = append(outputArr, item)
				delete(numFrequencies, item)
			}
		}
	}

	notExistedElements := make([]int, 0, len(numFrequencies))

	for key, frequencies := range numFrequencies {
		for index := 1; index <= frequencies; index++ {
			notExistedElements = append(notExistedElements, key)

		}
	}

	slices.SortFunc(notExistedElements, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	outputArr = append(outputArr, notExistedElements...)

	return outputArr
}

func relativeSortArrayUsingArr(arr1 []int, arr2 []int) []int {
	lenArr1 := len(arr1)
	outputArr := make([]int, 0, lenArr1)
	frequencies := make([]int, 1001)

	// Step 1: Count frequency each of element in arr1
	for _, element := range arr1 {
		frequencies[element]++
	}

	// Step 2: Add element to output arr that is followed by order of arr2
	for _, element := range arr2 {
		for frequencies[element] > 0 {
			outputArr = append(outputArr, element)
			frequencies[element]--
		}
	}

	for index := range frequencies {
		for frequencies[index] > 0 {
			outputArr = append(outputArr, index)
			frequencies[index]--
		}
	}

	return outputArr
}
