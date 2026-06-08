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
			RunBinarySearch()

		case 2:

		case 3:
			RunIsPerfectSquare()

		case 4:

		case 5:

		case 6:

		case 7:

		case 8:

		case 0:
			fmt.Println("Exit...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}

}

func showMenu() {
	fmt.Println("===== Review 10.Graph =====")
	fmt.Println("0. Exit")
	fmt.Println("1. 704. Binary Search")
	fmt.Println("2. 374. Guess Number Higher or Lower")
	fmt.Println("3. 367. Valid Perfect Square")
	fmt.Println("4. 441. Arranging Coins")
	fmt.Println("5. 35. Search Insert Position")
	fmt.Println("6. 34. Find First and Last Position of Element in Sorted Array")
	fmt.Println("7. 540. Single Element in a Sorted Array")
	fmt.Println("8. 74. Search a 2D Matrix")
	fmt.Println("9. 2300. Successful Pairs of Spells and Potions")
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
