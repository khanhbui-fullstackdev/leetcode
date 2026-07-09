package main

import "fmt"

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
	removeElement(nums, val)
}

func removeElement(nums []int, val int) int {
	return 0
}
