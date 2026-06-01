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
			RunTwoRun()

		case 2:
			RunSumOfUnique()

		case 3:
			//RunCanConstructUsingMap()
			RunCanConstruct()

		case 4:
			RunFindLucky()

		case 5:
			RunCountBalls()
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
	fmt.Println("===== Review 7.Hash Table =====")
	fmt.Println("0. Exit")
	fmt.Println("1. Two Sum")
	fmt.Println("2. 1748. Sum of Unique Elements")
	fmt.Println("3. 383. Ransom Note")
	fmt.Println("4. 1394. Find Lucky Integer in an Array")
	fmt.Println("5. 11742. Maximum Number of Balls in a Box")
	fmt.Println("6. 771. Jewels and Stones")
	fmt.Println("7. 49. Group Anagrams")
	fmt.Println("8. 2405. Optimal Partition of String")
	fmt.Println("9. 146. LRU Cache")
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
