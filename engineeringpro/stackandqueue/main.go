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
			RunIsValidParentheses()

		case 2:
			RunBackSpaceCompare()

		case 3:
			RunPing()

		case 4:
			RunCountStudents()

		case 5:
			RunRemoveDuplicates()

		case 6:
			RunTimeRequiredToBuy()

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
	fmt.Println("1. 20. Valid Parentheses")
	fmt.Println("2. 844. Backspace String Compare")
	fmt.Println("3. 933. Number of Recent Calls")
	fmt.Println("4. 1700. Number of Students Unable to Eat Lunch")
	fmt.Println("5. 1047. Remove All Adjacent Duplicates In String")
	fmt.Println("6. 2073. Time Needed to Buy Tickets")
	fmt.Println("7. ")
	fmt.Println("8. ")
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
