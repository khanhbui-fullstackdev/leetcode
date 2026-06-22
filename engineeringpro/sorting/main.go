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
			RunTargetIndicates()

		case 2:
			RunSmallerNumbersThanCurrent()

		case 3:

		case 4:

		case 5:

		case 6:

		case 7:

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
	fmt.Println("===== Review 4.Sorting =====")
	fmt.Println("0. Exit")
	fmt.Println("1. 2089. Find Target Indices After Sorting Array")
	fmt.Println("2. 1365. How Many Numbers Are Smaller Than the Current Number")
	fmt.Println("3. 1122. Relative Sort Array")
	fmt.Println("4. 88. Merge Sorted Array")
	fmt.Println("5. 1051. Height Checker")
	fmt.Println("6. 2144. Minimum Cost of Buying Candies With Discount")
	fmt.Println("7. 2357. Make Array Zero by Subtracting Equal Amounts")
	fmt.Println("8. 1637. Widest Vertical Area Between Two Points Containing No Points")
	fmt.Println("9. 2545. Sort the Students by Their Kth Score")
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
