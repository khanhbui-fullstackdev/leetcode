package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		showMenu()
		choice := readInput()

		switch choice {
		case 1:
			RunMaxProfit()

		case 2:
			RunRemoveElement()

		case 3:
			RunPivotIndex()

		case 4:
			RunFindDisappearedNumbers()

		case 5:
			RunSumRange()

		case 6:
			RunCanPlaceFlowers()

		case 7:
			RunProductExceptSelf()

		case 8:
			RunGeneratePascalTriangle()

		case 10:
			RunValidSudoku()

		case 0:
			fmt.Println("Exit...")

			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}

func showMenu() {
	fmt.Println("===== Review 1.Array =====")
	fmt.Println("0. Exit")
	fmt.Println("1. 121. Best Time to Buy and Sell Stock")
	fmt.Println("2. 27. Remove Element")
	fmt.Println("3. 724. Find Pivot Index")
	fmt.Println("4. 448. Find All Numbers Disappeared in an Array")
	fmt.Println("5. 303. Range Sum Query - Immutable")
	fmt.Println("6. 605. Can Place Flowers")
	fmt.Println("7. 238. Product of Array Except Self")
	fmt.Println("8. 118. Pascal's Triangle")
	fmt.Println("9. 122. Best Time to Buy and Sell Stock II")
	fmt.Println("10. 36. Valid Sudoku")
	fmt.Print("Enter your choice: ")
}

func readInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return -1 // invalid
	}

	return num
}
