package main

import (
	"cmp"
	"fmt"
	"math"
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

func RunThirdMax() {
	nums := []int{3, 2, 1}
	theThirdMax := thirdMaxUsing3Loops(nums)
	fmt.Printf("Nums:%v => Third max using 3 for:%d", nums, theThirdMax)
	fmt.Println()

	theThirdMax = thirdMaxOnly1Loop(nums)
	fmt.Printf("Nums:%v => Third max using one loop:%d", nums, theThirdMax)
	fmt.Println()

	theThirdMax = thirdMaxUsingSort(nums)
	fmt.Printf("Nums:%v => Third max using sort:%d", nums, theThirdMax)
	fmt.Println()

	nums = []int{1, -2147483648, 2}
	theThirdMax = thirdMaxUsing3Loops(nums)
	fmt.Printf("Nums:%v => Third max using 3 for:%d", nums, theThirdMax)
	fmt.Println()

	theThirdMax = thirdMaxOnly1Loop(nums)
	fmt.Printf("Nums:%v => Third max using one loop:%d", nums, theThirdMax)
	fmt.Println()

	theThirdMax = thirdMaxUsingSort(nums)
	fmt.Printf("Nums:%v => Third max using sort:%d", nums, theThirdMax)
	fmt.Println()

	nums = []int{1, 2}
	theThirdMax = thirdMaxUsing3Loops(nums)
	fmt.Printf("Nums:%v => Third max using 3 for:%d", nums, theThirdMax)
	fmt.Println()

	theThirdMax = thirdMaxOnly1Loop(nums)
	fmt.Printf("Nums:%v => Third max using one loop:%d", nums, theThirdMax)
	fmt.Println()

	theThirdMax = thirdMaxUsingSort(nums)
	fmt.Printf("Nums:%v => Third max using sort:%d", nums, theThirdMax)
	fmt.Println()

	nums = []int{1, 1, 2}
	theThirdMax = thirdMaxUsing3Loops(nums)
	fmt.Printf("Nums:%v => Third max using 3 for:%d", nums, theThirdMax)
	fmt.Println()

	theThirdMax = thirdMaxOnly1Loop(nums)
	fmt.Printf("Nums:%v => Third max using one loop:%d", nums, theThirdMax)
	fmt.Println()

	theThirdMax = thirdMaxUsingSort(nums)
	fmt.Printf("Nums:%v => Third max using sort:%d", nums, theThirdMax)
	fmt.Println()
}

// First solution: we sort the array should be desc then we can loop through to find the third number
// Sort: O(nlogn)
func thirdMaxUsingSort(nums []int) int {
	fmt.Println("*** Third max using sort ***")
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	firstMax := nums[0]
	if len(nums) <= 2 {
		return firstMax
	}

	secondMax := math.MinInt32 - 1
	thirdMax := math.MinInt32 - 1
	for _, num := range nums {
		if secondMax < num && num < firstMax {
			secondMax = num
		}
	}

	for _, num := range nums {
		if num < secondMax && thirdMax < num {
			thirdMax = num
		}
	}

	// corner case: if the third max does not exist, return the maximum number
	if thirdMax < math.MinInt32 {
		thirdMax = firstMax
	}

	return thirdMax
}

// We will use 3 loops to find firstMax,secondMax and the thirdMax
// Time complexity = O(n) + O(n) + O(n) = O(3n) = O(n)
func thirdMaxUsing3Loops(nums []int) int {
	fmt.Println("*** Third max using 3 loops ***")
	firstMax := math.MinInt32 - 1
	// first loop to find the firstMax
	for _, num := range nums {
		if num > firstMax {
			firstMax = num
		}
	}
	if len(nums) <= 2 {
		return firstMax
	}

	secondMax := math.MinInt32 - 1
	thirdMax := math.MinInt32 - 1
	//second loop to find the secondMax
	for _, num := range nums {
		if num > secondMax && num < firstMax {
			secondMax = num
		}
	}

	//third loop to find the thirdMax
	for _, num := range nums {
		if num > thirdMax && num < secondMax {
			thirdMax = num
		}
	}

	//ad-hoc case, if the third maximum does not exist, return the first maximum instead
	if thirdMax < math.MinInt32 {
		thirdMax = firstMax
	}

	return thirdMax
}

// We only iterate one time O(n)
func thirdMaxOnly1Loop(nums []int) int {
	minimumNum := math.MinInt32 - 1
	firstMax, secondMax, thirdMax := minimumNum, minimumNum, minimumNum

	for _, num := range nums {
		if num > firstMax {
			thirdMax = secondMax
			secondMax = firstMax
			firstMax = num
		}
		if num < firstMax && num > secondMax {
			thirdMax = secondMax
			secondMax = num
		}
		if num < secondMax && num > thirdMax {
			thirdMax = num
		}
	}

	//adhoc: if third distinct maximum does not exist, we return the first max instead
	if len(nums) <= 2 || thirdMax < math.MinInt32 {
		return firstMax
	}

	return thirdMax
}

func RunRunningSum() {
	nums := []int{1, 2, 3, 4}
	sumofNums := runningSum(nums)
	fmt.Printf("Nums:%v => Running sum of nums:%d", nums, sumofNums)
	fmt.Println()

	nums = []int{1, 1, 1, 1, 1}
	sumofNums = runningSum(nums)
	fmt.Printf("Nums:%v => Running sum of nums:%d", nums, sumofNums)
	fmt.Println()

	nums = []int{3, 1, 2, 10, 1}
	sumofNums = runningSum(nums)
	fmt.Printf("Nums:%v => Running sum of nums:%d", nums, sumofNums)
	fmt.Println()
}

func runningSum(nums []int) []int {
	for index, num := range nums {
		if index == 0 {
			continue
		}
		preNum := nums[index-1]
		sum := num + preNum
		nums[index] = sum
	}
	return nums
}

func RunGetConcatenation() {
	nums := []int{1, 2, 1}
	ans := getConcatenation(nums)
	fmt.Printf("Nums:%v => Ans:%v", nums, ans)
	fmt.Println()

	nums = []int{1, 3, 2, 1}
	ans = getConcatenation(nums)
	fmt.Printf("Nums:%v => Ans:%v", nums, ans)
	fmt.Println()
}

func getConcatenation(nums []int) []int {
	numsLength := len(nums)
	ans := make([]int, numsLength*2)
	for index, num := range nums {
		ans[index] = num
		ans[index+numsLength] = num
	}

	return ans
}

func RunRemoveDuplicates() {
	nums := []int{1, 1, 2}
	uniqueElements := removeDuplicates(nums)
	fmt.Printf("Nums:%v => Unique elements:%d", nums, uniqueElements)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	uniqueElements = removeDuplicates(nums)
	fmt.Printf("Nums:%v => Unique elements:%d", nums, uniqueElements)
}

func removeDuplicates(nums []int) int {
	return 0
}
