package main

import (
	"cmp"
	"fmt"
	"leetcode/neetcode/hashing/models"
	"slices"
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

func RunIntersection() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	intersectedArr := intersection(nums1, nums2)
	fmt.Printf("Intersected arr:%v", intersectedArr)
	fmt.Println()

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}

	intersectedArr = intersection(nums1, nums2)
	fmt.Printf("Intersected arr:%v", intersectedArr)
	fmt.Println()

	nums1 = []int{1, 2, 2, 1}
	nums2 = []int{2}

	intersectedArr = intersection(nums1, nums2)
	fmt.Printf("Intersected arr:%v", intersectedArr)
	fmt.Println()

}

/*
 */
func intersection(nums1 []int, nums2 []int) []int {
	arrInterSection := make([]int, 0, len(nums1))
	frequencyNum := make(map[int]int, len(nums1))

	for _, num := range nums1 {
		frequencyNum[num]++
	}

	for _, num := range nums2 {
		_, hasExisted := frequencyNum[num]
		if hasExisted {
			arrInterSection = append(arrInterSection, num)
			delete(frequencyNum, num)
		}
	}

	return arrInterSection
}

func RunIntersection2() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	intersectedArr := intersect(nums1, nums2)
	fmt.Printf("Intersected arr:%v", intersectedArr)
	fmt.Println()

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}

	intersectedArr = intersect(nums1, nums2)
	fmt.Printf("Intersected arr:%v", intersectedArr)
	fmt.Println()
}

func intersect(nums1 []int, nums2 []int) []int {
	intersection := make([]int, 0, len(nums1))
	numsFrequency := make(map[int]int, len(nums1))

	for _, num := range nums1 {
		numsFrequency[num]++
	}

	for _, num := range nums2 {
		frequency, hasExisted := numsFrequency[num]
		if hasExisted {
			frequency--
			if frequency < 0 {
				delete(numsFrequency, num)
			} else {
				intersection = append(intersection, num)
			}
			numsFrequency[num] = frequency
		}
	}

	return intersection
}

func RunTopKElement() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Printf("Nums:%v and k:%d => Top K frequenct:%v", nums, k, topKFrequent(nums, k))
	fmt.Println()
	fmt.Printf("Nums:%v and k:%d => Top K frequenct V2:%v", nums, k, topKFrequentV2(nums, k))
	fmt.Println()

	nums = []int{1}
	k = 1
	fmt.Printf("Nums:%v and k:%d => Top K frequenct:%v", nums, k, topKFrequent(nums, k))
	fmt.Println()
	fmt.Printf("Nums:%v and k:%d => Top K frequenct V2:%v", nums, k, topKFrequentV2(nums, k))
	fmt.Println()

	nums = []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
	k = 2
	fmt.Printf("Nums:%v and k:%d => Top K frequenct:%v", nums, k, topKFrequent(nums, k))
	fmt.Println()
	fmt.Printf("Nums:%v and k:%d => Top K frequenct V2:%v", nums, k, topKFrequentV2(nums, k))
	fmt.Println()

	nums = []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	k = 2
	fmt.Printf("Nums:%v and k:%d => Top K frequenct:%v", nums, k, topKFrequent(nums, k))
	fmt.Println()
	fmt.Printf("Nums:%v and k:%d => Top K frequenct V2:%v", nums, k, topKFrequentV2(nums, k))
	fmt.Println()
}

func topKFrequent(nums []int, k int) []int {
	mapFrequencies := make(map[int]int, len(nums))
	mostFrequentElements := make([]int, 0, len(nums))

	for _, num := range nums {
		mapFrequencies[num]++
	}

	frequentNums := make([]models.FrequentNum, 0, len(nums))
	for num, frequency := range mapFrequencies {
		frequenNum := models.NewFrequentNum(num, frequency)
		frequentNums = append(frequentNums, frequenNum)
	}

	clear(mapFrequencies)

	// sort by frequency desc
	slices.SortFunc(frequentNums, func(a, b models.FrequentNum) int {
		return b.Count - a.Count
	})

	for index, frequentNum := range frequentNums {
		if index == k {
			break
		}
		mostFrequentElements = append(mostFrequentElements, frequentNum.Num)
	}

	return mostFrequentElements
}

/*
          0 1 2 3 4 5
	nums [1,1,1,2,2,2], k = 2

	countingAppearance = [{1:3},{2:3}]

	loop through an dictionary in reversed order

	          0 1 2   3    4   5  6
	bucket = [      [1,2]          ] // 2 array dimenstion

          0 1 2 3 4 5
	nums [2,2,2,2,2,2], k = 2
	countingAppearance = [{2:3}]

	loop through an dictionary in reversed order
	          0 1   2   3    4   5  6
	bucket = [                     [2]] // 2 array dimenstion

	      0 1 2 3 4 5
	nums [1,1,1,2,2,5], k = 2

	countingAppearance = [{1:3},{2:2},{5:1}]

	loop through an dictionary in reversed order

	          0   1     2     3    4   5   6
	bucket = [   [5]   [2]   [1]           ] // 2 array dimenstion
	**/

func topKFrequentV2(nums []int, k int) []int {
	numsLength := len(nums)
	countingAppearance := make(map[int]int, numsLength)
	for _, num := range nums {
		countingAppearance[num]++
	}

	// index represents the appearing numbers (frequent number)
	bucket := make([][]int, numsLength+1)
	bucketLength := len(bucket)

	for num, frequentNum := range countingAppearance {
		bucket[frequentNum] = append(bucket[frequentNum], num)
	}
	mostFrequentElements := make([]int, 0, numsLength)

	for index := bucketLength - 1; index >= 0; index-- {
		arr := bucket[index]
		if len(arr) == 0 {
			continue
		}
		if k == 0 {
			break
		}
		for _, num := range arr {
			mostFrequentElements = append(mostFrequentElements, num)
			k--
		}
	}
	return mostFrequentElements
}

func RunLongestConsecutive() {
	nums := []int{100, 4, 200, 1, 3, 2}
	longestElements := longestConsecutiveSortingSolution(nums)
	fmt.Printf("Longest consecutive elements:%v", longestElements)
	fmt.Println()
	fmt.Printf("Longest consecutive elements V2:%v", longestConsecutiveUsingHashMap(nums))
	fmt.Println()

	// nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	// longestElements = longestConsecutiveSortingSolution(nums)
	// fmt.Printf("Longest consecutive elements:%v", longestElements)
	// fmt.Println()
	// fmt.Printf("Longest consecutive elements V2:%v", longestConsecutiveUsingHashMap(nums))
	// fmt.Println()

	// nums = []int{1, 0, 1, 2}
	// longestElements = longestConsecutiveSortingSolution(nums)
	// fmt.Printf("Longest consecutive elements:%v", longestElements)
	// fmt.Println()
	// fmt.Printf("Longest consecutive elements V2:%v", longestConsecutiveUsingHashMap(nums))
	// fmt.Println()

	// nums = []int{1, 2, 3, 10, 11, 12, 13}
	// longestElements = longestConsecutiveSortingSolution(nums)
	// fmt.Printf("Longest consecutive elements:%v", longestElements)
	// fmt.Println()
	// fmt.Printf("Longest consecutive elements V2:%v", longestConsecutiveUsingHashMap(nums))
	// fmt.Println()
}

func longestConsecutiveSortingSolution(nums []int) int {
	numLength := len(nums)
	if numLength <= 1 {
		return numLength
	}

	// O(nlogn)
	// nums = [1, 2, 3, 4, 100, 200]
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	count := 1
	maxCount := 0

	for index, num := range nums {
		nextIndex := index + 1
		if nextIndex == numLength {
			break
		}
		nextNum := nums[nextIndex]
		if nextNum-num == 1 {
			count++
		} else if nextNum == num {
			continue
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 1
		}
	}

	if count > maxCount {
		maxCount = count
	}

	return maxCount
}

func longestConsecutiveUsingHashMap(nums []int) int {
	numsLength := len(nums)

	if numsLength <= 0 {
		return numsLength
	}

	hashSet := make(map[int]bool)
	for _, num := range nums {
		hashSet[num] = true
	}
	maxLength := 0

	for key, _ := range hashSet {
		leftNeighbor := key - 1
		_, existed := hashSet[leftNeighbor]
		if !existed {
			// 1. Tìm thấy đầu chuỗi, khởi tạo chuỗi hiện tại dài 1
			currNum := key
			currLength := 1

			// 2. Vòng lặp đếm tiếp các số liên tiếp đằng sau: key+1, key+2, key+3...
			for hashSet[currNum+1] {
				currNum++
				currLength++
			}
			// 3. Cập nhật kỷ lục max
			if currLength > maxLength {
				maxLength = currLength
			}
		}
	}

	return maxLength
}
