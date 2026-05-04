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
			RunIsSameTree()

		case 2:
			RunIsSymmetric()

		case 3:
			RunMaxDepth()

		case 4:
			RunHasPathSumI()

		case 5:
			RunPathSum()

		case 6:
			RunInvertTree()

		case 7:
			RunIsBalanced()

		case 8:
			RunBinaryTreePaths()

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
	fmt.Println("1. 100. Same Tree")
	fmt.Println("2. 101. Symmetric Tree")
	fmt.Println("3. 559. Maximum Depth of N-ary Tree")
	fmt.Println("4. 112. Path Sum")
	fmt.Println("5. 113. Path Sum II")
	fmt.Println("6. 226. Invert Binary Tree")
	fmt.Println("7. 110. Balanced Binary Tree")
	fmt.Println("8. 257. Binary Tree Paths")
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
