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
			RunContainsDuplicate()

		case 2:

		case 3:
			RunIntersection()

		case 4:
			RunIntersection2()

		case 5:
			RunFirstUniqChar()

		case 6:

		case 7:
			RunTopKElement()

		case 8:
			RunLongestConsecutive()

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
	fmt.Println("===== 4.Hashing =====")
	fmt.Println("0. Exit")
	fmt.Println("1. 217. Contains Duplicate")
	fmt.Println("2. 242. Valid Anagram")
	fmt.Println("3. 349. Intersection of Two Arrays")
	fmt.Println("4. 350. Intersection of Two Arrays II")
	fmt.Println("5. 387. First Unique Character in a String")
	fmt.Println("6. ")
	fmt.Println("7. 347. Top K Frequent Elements")
	fmt.Println("8. 128. Longest Consecutive Sequence")
	fmt.Println("9. ")
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
