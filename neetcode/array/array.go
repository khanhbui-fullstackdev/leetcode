package main

import (
	"cmp"
	"fmt"
	"slices"
)

func RunTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("Target = %d => Indices of two numbers:%v", target, twoSum(nums, target))
	fmt.Println()

	nums = []int{3, 2, 4}
	target = 6
	fmt.Printf("Target = %d => Indices of two numbers:%v", target, twoSum(nums, target))
	fmt.Println()

	nums = []int{3, 3}
	target = 6
	fmt.Printf("Target = %d => Indices of two numbers:%v", target, twoSum(nums, target))
	fmt.Println()
}

func twoSum(nums []int, target int) []int {
	lengthNums := len(nums)

	indicesOfTwoNumbers := make([]int, 2)
	mapNums := make(map[int]int, lengthNums)

	for index, num := range nums {
		partner := target - num
		// check the partner has checkedin/existed in the system or not
		preIndex, existed := mapNums[partner]
		if existed {
			// if the parter already existed in system
			// then we append into array
			indicesOfTwoNumbers[0] = preIndex
			indicesOfTwoNumbers[1] = index
			break
		} else {
			mapNums[num] = index
		}
	}

	return indicesOfTwoNumbers
}

func RunRemoveElement() {
	nums := []int{3, 2, 2, 3}
	val := 3
	remainingElements := removeElementV1(nums, val)
	fmt.Printf("Remaining elemements:%d", remainingElements)
	fmt.Println()

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	remainingElements = removeElementV1(nums, val)
	fmt.Printf("Remaining elemements:%d", remainingElements)
	fmt.Println()
}

// Two pointers Ti's solution
// Using 2 pointers i starting from 0; and j starting from len(nums)-1
func removeElementV1(nums []int, val int) int {
	pointerI := 0
	pointerJ := len(nums) - 1

	for pointerI <= pointerJ {
		if nums[pointerI] != val {
			pointerI++
		} else {
			temp := nums[pointerJ]
			nums[pointerI] = nums[pointerJ]
			nums[pointerJ] = temp

			pointerJ--
		}
	}

	return pointerI
}

func RunMerge() {
	nums1 := []int{1, 2, 8, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	mergeUsingCopy(nums1, m, nums2, n)
	mergeUsingExtraMemory(nums1, m, nums2, n)
	merge(nums1, m, nums2, n)

	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	mergeUsingCopy(nums1, m, nums2, n)
	mergeUsingExtraMemory(nums1, m, nums2, n)
	merge(nums1, m, nums2, n)

	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	mergeUsingCopy(nums1, m, nums2, n)
	mergeUsingExtraMemory(nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
}

/*
	  This approach is using copy's api golang
	  Props: Easy to implement
		Step 1: Copy all element of array nums2 to nums1
		Step 2: Sort nums1

=> Time complexity O (nlogn)
*/
func mergeUsingCopy(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("*** Merge using copy ***")

	// target: nums1
	copy(nums1[m:], nums2)

	fmt.Printf("After executing -> Nums1:%v", nums1)

	// Sorting
	slices.SortFunc(nums1, func(a, b int) int {
		return cmp.Compare(a, b)
	})
	fmt.Printf("\n After sorting -> Nums1:%v", nums1)
	fmt.Println()
}

/*
Create new extra array called mergedNums
then we will compare value between nums1[idx] && nums2[jdx].
We simply check nums1[idx]<=num2s[jdx] // Brute Force technique -> add nums1[idx] into arr mergedNums
Otherwise, we -> add num2s[jdx]  into arr mergedNums

Brute Force technique
Add the remaining items from nums1 into extra array. Similarly to num2s
*/
func mergeUsingExtraMemory(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("*** Merge using extra memory ***")
	idx := 0
	jdx := 0
	mergedSortedNums := make([]int, 0, len(nums1))
	for idx < m && jdx < n {
		num1 := nums1[idx]
		num2 := nums2[jdx]

		if num1 <= num2 {
			mergedSortedNums = append(mergedSortedNums, num1)
			idx++
		} else {
			mergedSortedNums = append(mergedSortedNums, num2)
			jdx++
		}
	}

	for idx < m {
		mergedSortedNums = append(mergedSortedNums, nums1[idx])
		idx++
	}

	for jdx < n {
		mergedSortedNums = append(mergedSortedNums, nums2[jdx])
		jdx++
	}

	copy(nums1, mergedSortedNums)
	fmt.Println("Merged nums:", nums1)
	fmt.Println()
}

/*
Using 3 pointers
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("*** Merge using 3 pointers ***")
	updatedIndex := len(nums1) - 1

	// start to loop from the end
	idx := m - 1
	jdx := n - 1

	// becasue we want to copy all of value/element from nums2 to nums1, so that we need to loop while j>=0
	for jdx >= 0 {
		if idx >= 0 && nums1[idx] >= nums2[jdx] {
			// pick the higher value num1
			nums1[updatedIndex] = nums1[idx]
			idx--
		} else {
			nums1[updatedIndex] = nums2[jdx]
			jdx--
		}
		updatedIndex--
	}
	fmt.Println("Merged nums:", nums1)
	fmt.Println()
}
