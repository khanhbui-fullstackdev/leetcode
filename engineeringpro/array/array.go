package main

import (
	"fmt"

	"leetcode/engineeringpro/array/models"
)

func RunMaxProfit() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Prices:%v -> Max profit:%d", prices1, maxProfit(prices1))

	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Printf("\nPrices:%v -> Max profit:%d", prices2, maxProfit(prices2))
	fmt.Println()
}

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := 10000

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		currentProfit := price - minPrice
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
	}

	return maxProfit
}

func RunRemoveElement() {
	nums1 := []int{3, 3, 3, 1}
	val1 := 3
	res1 := removeElement(nums1, val1)
	fmt.Printf("Arr:%v with val %d -> %d", nums1, val1, res1)

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	res2 := removeElement(nums2, val2)
	fmt.Printf("\nArr:%v with val %d -> %d", nums2, val2, res2)
	fmt.Println()
}

func removeElement(nums []int, val int) int {
	writerIdx := 0

	for _, num := range nums {
		if num != val {
			nums[writerIdx] = num
			writerIdx++
		}
	}

	return writerIdx
}

func RunPivotIndex() {
	nums1 := []int{1, 7, 3, 6, 5, 6}
	fmt.Printf("Pivot index of nums=%v is %d", nums1, pivotIndex(nums1))
	fmt.Println()

	nums2 := []int{1, 2, 3}
	fmt.Printf("Pivot index of nums=%v is %d", nums2, pivotIndex(nums2))
	fmt.Println()

	nums3 := []int{2, 1, -1}
	fmt.Printf("Pivot index of nums=%v is %d", nums3, pivotIndex(nums3))
	fmt.Println()
}

func pivotIndex(nums []int) int {
	preSum := 0
	for _, num := range nums {
		preSum += num
	}

	leftSum := 0
	pivotIndex := -1

	for index, num := range nums {
		if index > 0 {
			leftSum += nums[index-1]
		}
		rightSum := preSum - num - leftSum
		if rightSum == leftSum {
			pivotIndex = index
		}
	}

	return pivotIndex
}

func RunFindDisappearedNumbers() {
	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}

	disappearedNumbers := findDisappearedNumbers(nums1)
	fmt.Printf("\n Disappeared numbers:%v", disappearedNumbers)
	fmt.Println()
}

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		pos := absInt(num) - 1
		if nums[pos] > 0 {
			nums[pos] = -nums[pos]
		}
	}

	fmt.Printf("\n Nums:%v", nums)

	var disappearedNumbers []int
	for index, num := range nums {
		if num > 0 {
			disappearedNumbers = append(disappearedNumbers, index+1)
		}
	}

	return disappearedNumbers
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func RunSumRange() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	numArr := models.Constructor(nums)

	fmt.Printf(" Pre sums:%v", numArr)
	fmt.Println()

	fmt.Printf("\n Sum range(0, 2) = %d", numArr.SumRange(0, 2))
	fmt.Printf("\n Sum range(2, 5) = %d", numArr.SumRange(2, 5))
	fmt.Printf("\n Sum range(0, 5) = %d", numArr.SumRange(0, 5))
	fmt.Println()
}

func RunCanPlaceFlowers() {
	flowerbeds1 := []int{1, 0, 0, 0, 1}
	n1 := 1

	fmt.Printf("Can place flowers:%v", canPlaceFlowers(flowerbeds1, n1))
	fmt.Println()

	flowerbeds2 := []int{1, 0, 0, 0, 1}
	n2 := 2
	fmt.Printf("Can place flowers:%v", canPlaceFlowers(flowerbeds2, n2))
	fmt.Println()
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for index, flower := range flowerbed {
		if n == 0 {
			return true
		}
		if flower == 0 {
			leftValid := index == 0 || flowerbed[index-1] == 0
			rightValid := index == len(flowerbed)-1 || flowerbed[index+1] == 0
			if leftValid && rightValid {
				flowerbed[index] = 1
				n--
				index += 2
			}
		}
	}
	return false
}

func RunProductExceptSelf() {
	nums := []int{1, 2, 3, 4}

	array := productExceptSelf(nums)
	fmt.Printf("Array:%v", array)
	fmt.Println()
}

func productExceptSelf(nums []int) []int {
	answers := make([]int, len(nums))

	left := 1
	for index, num := range nums {
		answers[index] = left
		left *= num
	}
	fmt.Println("Left product:", answers)

	right := 1
	for index := len(nums) - 1; index >= 0; index-- {
		answers[index] = answers[index] * right
		fmt.Printf("\n answers[%d]:%d", index, answers[index])
		right = right * nums[index]
		fmt.Printf("\n right:%d", right)
	}

	return answers
}

func RunValidSudoku() {
	// Valid case
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	isValidSudoku1 := isValidSudokuUsingMap(board1)
	fmt.Printf("Is valid Sudoku: %v", isValidSudoku1)
	fmt.Println()

	// Invalid case
	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	isValidSudoku2 := isValidSudokuUsingMap(board2)
	fmt.Printf("Is valid Sudoku: %v", isValidSudoku2)
	fmt.Println()
}

func isValidSudokuUsingMap(board [][]byte) bool {
	boardLength := len(board)
	fmt.Println("Board length:", boardLength)

	sudokuMap := make(map[string]bool, boardLength)

	for rowIndex := range boardLength {
		for colIndex := range boardLength {
			asciiNumber := board[rowIndex][colIndex]
			if asciiNumber == '.' {
				continue
			}
			number := int(asciiNumber - '0')

			rowKey := fmt.Sprintf("\n %d has existed in row %d", number, rowIndex)
			colKey := fmt.Sprintf("\n %d has existed in col %d", number, colIndex)
			boxKey := fmt.Sprintf("\n %d has existed in box %d", number, rowIndex/3*3+colIndex/3)

			isValidRow, _ := sudokuMap[rowKey]
			isValidCol, _ := sudokuMap[colKey]
			isValidBox, _ := sudokuMap[boxKey]
			if isValidRow || isValidCol || isValidBox {
				return false
			}

			sudokuMap[rowKey] = true
			sudokuMap[colKey] = true
			sudokuMap[boxKey] = true
		}
	}

	fmt.Println("\n Sudoku board:", sudokuMap)
	return true
}

func RunGeneratePascalTriangle() {
	numRows := 5
	pascalTriangle := generate(numRows)

	fmt.Println("Pascal triangle:", pascalTriangle)
	fmt.Println()
}

/*
			1                         row = 0 -> pascalTriangle = [[1]]
		  1	  1                       row = 1 -> pascalTriangle= [[1], [1,1]]
        1   2   1                     row = 2 -> arr has 3 elements -> [1, pascalTriangle[1][0] +  pascalTriangle[1][1], 1] = arr[1,2,1]
      1   3   3   1					  row = 3 -> arr has 4 elements -> [1, pascalTriangle[2][0] +  pascalTriangle[2][1], pascalTriangle[2][1] + pascalTriangle[2][2], 1]
	1   4   6   4   1                 row = 4 -> arr has 5 elements -> [1, pascalTriangle[3][0] +  pascalTriangle[3][1], pascalTriangle[3][1] + pascalTriangle[3][2], pascalTriangle[3][2] + pascalTriangle[3][3]
*/

func generate(numRows int) [][]int {
	pascalTriangle := make([][]int, 0, numRows) //  len = 0, cap = 5

	for row := range numRows {
		// len
		arr := make([]int, row+1) //  len = 0, cap = 1  every row will contain n + 1 elements
		arr[0] = 1
		arr[len(arr)-1] = 1
		if row > 1 {
			for col := 1; col < row; col++ {
				a := pascalTriangle[row-1][col-1]
				b := pascalTriangle[row-1][col]
				fmt.Println("a:", a)
				fmt.Println("b:", b)
				arr[col] = a + b
				fmt.Println("Arr:", arr)
			}
		}
		pascalTriangle = append(pascalTriangle, arr)
	}

	return pascalTriangle
}
