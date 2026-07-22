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

		case 2:

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
	fmt.Println("===== Review 1.Array =====")
	fmt.Println("0. Exit")
	fmt.Println("1. 21. Merge Two Sorted Lists")
	fmt.Println("2. 83. Remove Duplicates from Sorted List")
	fmt.Println("3. 203. Remove Linked List Elements")
	fmt.Println("4. 206. Reverse Linked List")
	fmt.Println("5. 876. Middle of the Linked List")
	fmt.Println("6. 141. Linked List Cycle")
	fmt.Println("7. 19. Remove Nth Node From End of List")
	fmt.Println("8. 2. Add Two Numbers")
	fmt.Println("9. 138. Copy List with Random Pointer")
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
