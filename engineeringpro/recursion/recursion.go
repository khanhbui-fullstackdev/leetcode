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
			RunFib()

		case 2:
			RunPowerOfThree()

		case 3:
			RunPowerOfFour()

		case 4:
			RunIsHappy()

		case 5:

		case 6:

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
	fmt.Println("===== Review 3.Recursion =====")
	fmt.Println("0. Exit")
	fmt.Println("1. 509. Fibonacci Number")
	fmt.Println("2. 326. Power of Three")
	fmt.Println("3. 342. Power of Four")
	fmt.Println("4. 202. Happy Number")
	fmt.Println("5. 1837. Sum of Digits in Base K")
	fmt.Println("6. 504. Base 7")
	fmt.Println("7. 779. K-th Symbol in Grammar")
	fmt.Println("8. 1545. Find Kth Bit in Nth Binary String")
	fmt.Println("9. 169. Tower of Hanoi (Optional)")
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
