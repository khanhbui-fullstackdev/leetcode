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
			RunTwoSum()

		case 2:
			RunRemoveElement()

		case 3:
			RunRunningSum()

		case 4:
			RunMerge()

		case 5:
			RunGetConcatenation()

		case 6:
			RunRemoveDuplicates()

		case 7:
			RunThirdMax()

		case 8:

		case 9:

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
	fmt.Println("1. Two Sum")
	fmt.Println("2. 27. Remove Element")
	fmt.Println("3. 1480. Running Sum of 1d Array")
	fmt.Println("4. 88. Merge Sorted Array")
	fmt.Println("5. 1929. Concatenation of Array")
	fmt.Println("6. 26. Remove Duplicates from Sorted Array")
	fmt.Println("7. 414. Third Maximum Number")
	fmt.Println("8. 15. 3Sum")
	fmt.Println("9. 11. Container With Most Water")
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
