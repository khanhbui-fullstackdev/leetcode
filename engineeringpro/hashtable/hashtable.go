package main

import (
	"fmt"
)

func RunTwoRun() {
	nums := []int{2, 7, 11, 15}
	target := 9
	arr := twoSum(nums, target)
	fmt.Printf("Array Two Sum:%v", arr)
	fmt.Println()

	nums = []int{3, 2, 4}
	target = 6
	arr = twoSum(nums, target)
	fmt.Printf("Array Two Sum:%v", arr)
	fmt.Println()

	nums = []int{3, 3}
	target = 6
	arr = twoSum(nums, target)
	fmt.Printf("Array Two Sum:%v", arr)
	fmt.Println()
}

func twoSum(nums []int, target int) []int {
	numLength := len(nums)
	twoSumArr := make([]int, 0, numLength)
	checkedInMaps := make(map[int]int, numLength)

	for index, num := range nums {
		offset := target - num
		preIndex, existed := checkedInMaps[offset]
		if existed {
			twoSumArr = append(twoSumArr, preIndex, index)
		} else {
			checkedInMaps[num] = index
		}
	}
	clear(checkedInMaps)
	return twoSumArr
}

func RunSumOfUnique() {
	nums := []int{1, 2, 3, 2}
	totalUnique := sumOfUnique(nums)
	fmt.Printf("Sum of unique:%v", totalUnique)
	fmt.Println()

	nums = []int{1, 1, 1, 1, 1}
	totalUnique = sumOfUnique(nums)
	fmt.Printf("Sum of unique:%v", totalUnique)
	fmt.Println()

	nums = []int{1, 2, 3, 4, 5}
	totalUnique = sumOfUnique(nums)
	fmt.Printf("Sum of unique:%v", totalUnique)
	fmt.Println()
}

func sumOfUnique(nums []int) int {
	uniqueSum := 0
	lenNums := len(nums)
	frequencyNums := make(map[int]int, lenNums)
	// first count the frequency of nums
	for _, num := range nums {
		frequencyNums[num]++
	}

	for num, frequency := range frequencyNums {
		if frequency == 1 {
			uniqueSum += num
		}
	}

	clear(frequencyNums)
	return uniqueSum
}

func RunCanConstruct() {
	ransomNote := "a"
	magazine := "b"
	fmt.Printf("\nRansome note %s can be constructed from letters of managinze %s is %v", ransomNote, magazine, canConstruct(ransomNote, magazine))
	fmt.Println()

	// ransomNote = "aa"
	// magazine = "ab"
	// fmt.Printf("\nRansome note %s can be constructed from letters of managinze %s is %v", ransomNote, magazine, canConstruct(ransomNote, magazine))
	// fmt.Println()

	// ransomNote = "aa"
	// magazine = "aab"
	// fmt.Printf("\nRansome note %s can be constructed from letters of managinze %s is %v", ransomNote, magazine, canConstruct(ransomNote, magazine))
	// fmt.Println()
}

func RunCanConstructUsingMap() {
	ransomNote := "a"
	magazine := "b"
	fmt.Printf("\nRansome note %s can be constructed from letters of managinze %s is %v", ransomNote, magazine, canConstructUsingMap(ransomNote, magazine))

	ransomNote = "aa"
	magazine = "ab"
	fmt.Printf("\nRansome note %s can be constructed from letters of managinze %s is %v", ransomNote, magazine, canConstructUsingMap(ransomNote, magazine))

	ransomNote = "aa"
	magazine = "aab"
	fmt.Printf("\nRansome note %s can be constructed from letters of managinze %s is %v", ransomNote, magazine, canConstructUsingMap(ransomNote, magazine))

	ransomNote = "fihjjjjei"
	magazine = "hjibagacbhadfaefdjaeaebgi"
	fmt.Printf("\nRansome note %s can be constructed from letters of managinze %s is %v", ransomNote, magazine, canConstructUsingMap(ransomNote, magazine))
	fmt.Println()
}

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteLen := len(ransomNote)
	magazineLen := len(magazine)

	if ransomNoteLen > magazineLen {
		return false
	}

	var countCharFromManagazine [26]int
	for _, runeM := range magazine {
		charM := int(runeM - 'a')
		countCharFromManagazine[charM]++
	}
	fmt.Println(countCharFromManagazine)
	for _, runeR := range ransomNote {
		charR := int(runeR - 'a')
		countCharFromManagazine[charR]--
		if countCharFromManagazine[charR] < 0 {
			fmt.Println(countCharFromManagazine)
			return false
		}
	}

	return true
}

func canConstructUsingMap(ransomNote string, magazine string) bool {
	ransomNoteLen := len(ransomNote)
	magazineLen := len(magazine)
	if ransomNoteLen > magazineLen {
		return false
	}
	sourceMagazine := make(map[int]int, magazineLen)
	for _, runeM := range magazine {
		sourceMagazine[int(runeM)]++
	}

	for _, runeR := range ransomNote {
		frequency, existed := sourceMagazine[int(runeR)]
		if existed {
			frequency--
			sourceMagazine[int(runeR)] = frequency
			if frequency < 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func RunFindLucky() {
	arr := []int{2, 2, 3, 4}
	fmt.Printf("\nLucky number:%d", findLucky(arr))

	arr = []int{1, 2, 2, 3, 3, 3}
	fmt.Printf("\nLucky number:%d", findLucky(arr))

	arr = []int{2, 2, 2, 3, 3}
	fmt.Printf("\nLucky number:%d", findLucky(arr))
	fmt.Println()
}

func findLucky(arr []int) int {
	arrLength := len(arr)
	mapArr := make(map[int]int, arrLength)

	for _, element := range arr {
		mapArr[element]++
	}

	largestLuckyNumber := -1
	for num, frequency := range mapArr {
		if frequency == num {
			if num > largestLuckyNumber {
				largestLuckyNumber = num
			}
		}
	}
	return largestLuckyNumber
}

func RunCountBalls() {
	lowLimit := 1
	highLimit := 10
	balls := countBalls(lowLimit, highLimit)
	fmt.Printf("\nBalls:%d", balls)

	lowLimit = 5
	highLimit = 15
	balls = countBalls(lowLimit, highLimit)
	fmt.Printf("\nBalls:%d", balls)
	fmt.Println()

	lowLimit = 19
	highLimit = 28
	balls = countBalls(lowLimit, highLimit)
	fmt.Printf("\nBalls:%d", balls)
	fmt.Println()
}

func countBalls(lowLimit int, highLimit int) int {
	ballNumbered := highLimit - lowLimit + 1
	boxMap := make(map[int]int, ballNumbered)

	for index := lowLimit; index <= highLimit; index++ {
		if index < 10 {
			boxMap[index]++
		} else {
			numberedBox := calculateNumberedBox(index, 0)
			boxMap[numberedBox]++
		}
	}

	numberOfMostBalls := 0
	for _, frequency := range boxMap {
		if numberOfMostBalls < frequency {
			numberOfMostBalls = frequency
		}
	}
	clear(boxMap)
	return numberOfMostBalls
}

func calculateNumberedBox(num int, remainer int) int {
	if num == 0 {
		return remainer
	}
	remainer += num % 10
	num = num / 10
	return calculateNumberedBox(num, remainer)
}
