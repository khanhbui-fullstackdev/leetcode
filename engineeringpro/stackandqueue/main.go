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
	fmt.Println("1. 20. Valid Parentheses")
	fmt.Println("2. 844. Backspace String Compare")
	fmt.Println("3. 933. Number of Recent Calls")
	fmt.Println("4. ")
	fmt.Println("5. ")
	fmt.Println("6. ")
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
