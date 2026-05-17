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
			RunMiddleNode()

		case 2:
			RunDeleteDuplicates()

		case 3:
			RunRemoveElement()

		case 4:
			RunGetIntersectionNode()

		case 5:
			RunReverseList()

		case 6:
			RunMergeTwoLists()

		case 7:

		case 8:

		case 9:

		case 10:

		case 11:

		case 0:
			fmt.Println("Exit...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}

}

func showMenu() {
	fmt.Println("===== Review 8.Linked List =====")
	fmt.Println("0. Exit")
	fmt.Println("1. 876. Middle of the Linked List")
	fmt.Println("2. 83. Remove Duplicates from Sorted List")
	fmt.Println("3. 203. Remove Linked List Elements")
	fmt.Println("4. 160. Intersection of Two Linked Lists")
	fmt.Println("5. 206. Reverse Linked List")
	fmt.Println("6. 21. Merge Two Sorted Lists")
	fmt.Println("7. ")
	fmt.Println("8. ")
	fmt.Println("9. ")
	fmt.Println("10. ")
	fmt.Println("11. ")
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
