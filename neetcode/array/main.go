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

		case 3:

		case 4:

		case 5:

		case 6:

		case 7:

		case 8:

		case 10:

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
	fmt.Println("7. 15. 3Sum")
	fmt.Println("8. 11. Container With Most Water")
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
