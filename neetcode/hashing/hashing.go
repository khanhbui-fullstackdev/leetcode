package main

import (
	"fmt"
	"leetcode/neetcode/hashing/models"
)

func RunFirstUniqChar() {
	s := "leetcode"
	firstNoneRepeatingChar := firstUniqChar(s)
	fmt.Printf("First non-repeating character:%d", firstNoneRepeatingChar)
	fmt.Println()

	s = "loveleetcode"
	firstNoneRepeatingChar = firstUniqChar(s)
	fmt.Printf("First non-repeating character:%d", firstNoneRepeatingChar)
	fmt.Println()

	s = "aabb"
	firstNoneRepeatingChar = firstUniqChar(s)
	fmt.Printf("First non-repeating character:%d", firstNoneRepeatingChar)
	fmt.Println()
}

func firstUniqChar(s string) int {
	sLength := len(s)
	charMap := make(map[rune]models.CharFrequency, sLength)
	for index, runeS := range s {
		updatedCharFrequency, hasExisted := charMap[runeS]
		if !hasExisted {
			indexes := []int{index}
			newCharFrequency := models.CharFrequency{Count: 1, Indexes: indexes}
			charMap[runeS] = newCharFrequency
		} else {
			updatedCharFrequency.Count++
			updatedCharFrequency.Indexes = append(updatedCharFrequency.Indexes, index)
			charMap[runeS] = updatedCharFrequency
		}
	}
	minIndex := 100000
	hasFirstUniqueChar := false
	for _, charFrequency := range charMap {
		if charFrequency.Count == 1 {
			hasFirstUniqueChar = true
			index := charFrequency.Indexes[0]
			if minIndex > index {
				minIndex = index
			}
		}
	}
	if !hasFirstUniqueChar {
		minIndex = -1
	}
	return minIndex
}

func RunContainsDuplicate() {
	nums := []int{1, 2, 3, 1}
	fmt.Printf("Is duplicated:%v", containsDuplicate(nums))
	fmt.Println()

	nums = []int{1, 2, 3, 4}
	fmt.Printf("Is duplicated:%v", containsDuplicate(nums))
	fmt.Println()

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Printf("Is duplicated:%v", containsDuplicate(nums))
	fmt.Println()
}

// Time complexity O(n)
// Space complexity O(n)
func containsDuplicate(nums []int) bool {
	numLength := len(nums)
	frequencyNums := make(map[int]int, numLength)
	for _, num := range nums {
		frequency, hasExisted := frequencyNums[num]
		if !hasExisted {
			frequencyNums[num] = 1
		} else {
			frequency++
			frequencyNums[num] = frequency
		}
		if frequency > 1 {
			return true
		}
	}
	return false
}
