package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func RunFib() {
	n := 2
	fmt.Printf("Fibonacci number fib(%d):%d", n, fibRecursion(n))
	fmt.Println()

	n = 3
	fmt.Printf("Fibonacci number fib(%d):%d", n, fibRecursion(n))
	fmt.Println()

	n = 4
	fmt.Printf("Fibonacci number fib(%d):%d", n, fibRecursion(n))
	fmt.Println()

	start := time.Now()
	n = 45 // it takes 4.5 seconds to complete
	fmt.Printf("Fibonacci number fib(%d):%d", n, fibRecursion(n))
	fmt.Println()
	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)

	// From n is greater or equal to 45, we must use dynamic programming
	start = time.Now()
	n = 100
	fibonacies := make([]int, n+1)
	//set default value is -1 // fibonacies = [-1 -1 -1 -1 -1,-1]
	fibonacies = setDefaultValueForSlice(fibonacies, -1)
	fibonacies[0] = 0
	fibonacies[1] = 1

	// fmt.Printf("Fibonacci number using dynamic programming topdown fib(%d):%d", n, fibDynamicProgrammingTopdown(n, fibonacies))
	// fmt.Println()
	// elapsed = time.Since(start)
	// fmt.Println("Elapsed:", elapsed)

	fmt.Printf("Fibonacci number using dynamic programming bottomup fib(%d):%d", n, fibDynamicProgrammingTopdown(n, fibonacies))
	fmt.Println()
	elapsed = time.Since(start)
	fmt.Println("Elapsed:", elapsed)
}

/*
Using recursion, everytime fun is called it will create f1(n-1) & f2(n-2)
=> Time complexity O(2^n)
=> Space complexity O(n)
*/
func fibRecursion(n int) int {
	if n < 2 {
		return n
	}
	f1 := fibRecursion(n - 1)
	f2 := fibRecursion(n - 2)
	return f1 + f2
}

/*
Using dynamic programming (top-down)
=> Time complexity O(n)
=> Space complexity O(n)
*/
func fibDynamicProgrammingTopdown(n int, fibonacies []int) int {
	if n < 2 {
		return n
	}

	if fibonacies[n] != -1 {
		return fibonacies[n]
	}
	f1 := fibDynamicProgrammingTopdown(n-1, fibonacies)
	f2 := fibDynamicProgrammingTopdown(n-2, fibonacies)
	fibonacies[n] = f1 + f2
	return fibonacies[n]
}

func setDefaultValueForSlice(nums []int, defaultValue int) []int {
	for index := range nums {
		if nums[index] == 0 {
			nums[index] = defaultValue
		}
	}
	return nums
}

func finDynamicProgrammingBottomup(n int, fibonacies []int) int {
	if n < 2 {
		return n
	}

	for index := 2; index <= n; index++ {
		fibonacies[index] = fibonacies[index-1] + fibonacies[index-2]
	}
	return fibonacies[n]
}

func RunPowerOfThree() {
	n := 27
	fmt.Printf("Is power of three:%v", isPowerOfThree(n))
	fmt.Println()

	n = 0
	fmt.Printf("Is power of three:%v", isPowerOfThree(n))
	fmt.Println()

	n = -1
	fmt.Printf("Is power of three:%v", isPowerOfThree(n))
	fmt.Println()
}

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 || n%3 != 0 {
		return false
	}
	return isPowerOfThree(n / 3)
}

func RunPowerOfFour() {
	n := 16
	fmt.Printf("Is power of four:%v", isPowerOfFour(n))
	fmt.Println()

	n = 5
	fmt.Printf("Is power of four:%v", isPowerOfFour(n))
	fmt.Println()

	n = 1
	fmt.Printf("Is power of four:%v", isPowerOfFour(n))
	fmt.Println()
}

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	} else if n == 0 || n%4 != 0 {
		return false
	}
	return isPowerOfFour(n / 4)
}

func RunIsHappy() {
	n := 19
	fmt.Printf("%d is a happy number:%v", n, isHappy(n))
	fmt.Println()

	n = 2
	fmt.Printf("%d is a happy number:%v", n, isHappy(n))
	fmt.Println()

	n = 199
	fmt.Printf("%d is a happy number:%v", n, isHappy(n))
	fmt.Println()

	n = 52
	fmt.Printf("%d is a happy number:%v", n, isHappy(n))
	fmt.Println()

	n = 2705
	fmt.Printf("%d is a happy number:%v", n, isHappy(n))
	fmt.Println()
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	squareOfDigits := make(map[int]bool, n)
	fmt.Println(squareOfDigits)
	return calculateSquareOfDigits(n, squareOfDigits)
}

func calculateSquareOfDigits(n int, squareOfDigits map[int]bool) bool {
	if n == 1 {
		return true
	}
	_, existed := squareOfDigits[n]
	if existed {
		return false
	}
	squareOfDigits[n] = true
	total := 0
	for n > 0 {
		digit := n % 10
		total = total + digit*digit
		n = n / 10
	}
	return calculateSquareOfDigits(total, squareOfDigits)
}

func RunSumBase() {
	n := 34
	k := 6
	sum := sumBaseV2(n, k)
	fmt.Printf("Sum of digits %d in Base K:%d", sum, k)
	fmt.Println()

	n = 10
	k = 10
	sum = sumBaseV2(n, k)
	fmt.Printf("Sum of digits %d in Base K:%d", sum, k)
	fmt.Println()
}

func sumBase(n int, k int) int {
	digits := make([]int, 0, n)
	calculateSumBaseV1(n, k, &digits)
	total := 0
	for _, digit := range digits {
		total += digit
	}
	return total
}

func calculateSumBaseV1(n, k int, digits *[]int) {
	if n == 0 {
		return
	}
	digit := (int)(n / k)
	remaining := (int)(n % k)
	*digits = append(*digits, remaining)
	calculateSumBaseV1(digit, k, digits)
}

func sumBaseV2(n int, k int) int {
	total := 0
	return calculateSumBaseV2(n, k, total)
}

func calculateSumBaseV2(n, k, total int) int {
	if n == 0 {
		return total
	}
	digit := (int)(n / k)
	remaining := (int)(n % k)
	total += remaining

	return calculateSumBaseV2(digit, k, total)
}

func RunConvertToBase7() {
	n := 100
	fmt.Printf("n = %d -> convert to base 7 = %s", n, convertToBase7(n))
	fmt.Println()

	n = -7
	fmt.Printf("n = %d -> convert to base 7:%s", n, convertToBase7(n))
	fmt.Println()
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}
	runes := make([]rune, 0, num)
	runes = calculateBase7(num, runes)

	var strBuilder strings.Builder
	strBuilder.Grow(len(runes))

	for index := len(runes) - 1; index >= 0; index-- {
		strBuilder.WriteRune(runes[index])
	}
	if isNegative {
		return "-" + strBuilder.String()
	}
	return strBuilder.String()
}

func calculateBase7(num int, runes []rune) []rune {
	if num == 0 {
		return runes
	}
	digit := num / 7
	remain := num % 7
	runes = append(runes, rune('0'+remain))
	return calculateBase7(digit, runes)
}

func RunKthGrammar() {
	n := 1
	k := 1
	fmt.Printf("At kth:%d and row(%d) is char:%d", k, n, kthGrammar(n, k))
	fmt.Println()

	n = 2
	k = 1
	fmt.Printf("At kth:%d and row(%d) is char:%d", k, n, kthGrammar(n, k))
	fmt.Println()

	n = 2
	k = 2
	fmt.Printf("At kth:%d and row(%d) is char:%d", k, n, kthGrammar(n, k))
	fmt.Println()

	n = 5
	k = 6
	fmt.Printf("At kth:%d and row(%d) is char:%d", k, n, kthGrammar(n, k))
	fmt.Println()

	n = 4
	k = 8
	fmt.Printf("At kth:%d and row(%d) is char:%d", k, n, kthGrammar(n, k))
	fmt.Println()
}

/*
	kthGrammar là bài đối xứng tịnh tiến
	row(1)=0
	  0 -> 10
	  combine = 10 = row(2)

	row(2)=10
      1 -> 10
	  0 -> 01
	  combine = 1001 = row(3)

	row(3)=1001
     1  -> 10
     0  -> 01
	 0  -> 01
	 1  -> 10
	 combine = 1001       0110           = row(4) => length = 8
   => row(4) = row(3)     invert(row(3))
	 Để xác định dc 1/2 & 1/2 sau ta phải xác định dc midIndex = sLength/2
     Để tính toán dc midIndex ta phải tính dc sLength
     sLength(N) = 2^(n-1) = sLength(5) = 2^(5-1) = 16

	 khi ta tính dc middleIndex ở hàng n
	 ta so sánh middleIndex với k

	 if middleIndex < k {
	    // k is located on the rightside of middleIndex
        // k nằm ở vị trí ĐẢO NGƯỢC BIT (1/2 phải)
		// theo tính chất đối xứng tịnh tiến k = 6(input) đang đối xứng k = 2
index	1234
		1001
        ||||
index   5678
		0110

		// dù k nằm ở vị trí 1/2 trc hay 1/2 sau thì ký tự ở vị trí thứ kth index đề nằm ở hàng trc đó n-1
		=> n = n - 1 = 4 - 1 = 3
		// đối với k > middleIndex => ta lấy k = k - middleIndex = 6 - 4 = 2 (index=2 đối  xứng index=6)
		// Ở 1/2 trái ta thấy index = 2 có ký tự 0, 1/2 phải index = 6 có ký tự 1
		=> 1 - 0 (1/2 trái) = 1 (1/2 phải) => 1 - (n,k)
	 }

k=2,n=3 {
         1234 (index)
row(3) = 1001
    	middleIndex = 2
		if middleIndex < k (2<2) ❌
}
		n = n - 1 = 2
		return kthGrammar(2,2)

k = 2, n = 2 {
         12 (index)
row(2) = 10
	middleIndex = 1
	if middleInex < k (1<2) ✅{
		n = n - 1 = 1
		k = k - middleIndex = 2 - 1 = 1
		return 1 - kthGrammar(1,1)
	}
}

k = 1, n = 1 {
	if n == 1 {
	  	return 0
	}
}
*/

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	sLength := int(math.Pow(2, float64(n-1)))
	middleIndex := sLength / 2
	if middleIndex < k {
		n = n - 1
		k = k - middleIndex
		return 1 - kthGrammar(n, k)
	}
	n = n - 1
	return kthGrammar(n, k)
}

func RunFindKthBit() {
	n := 3
	k := 1
	fmt.Printf("n = %d and k = %d => kthBit:%c", n, k, findKthBit(n, k))
	fmt.Println()

	n = 4
	k = 11
	fmt.Printf("n = %d and k = %d => kthBit:%c", n, k, findKthBit(n, k))
	fmt.Println()

	n = 3
	k = 5
	fmt.Printf("n = %d and k = %d => kthBit:%c", n, k, findKthBit(n, k))
	fmt.Println()

	n = 3
	k = 7
	fmt.Printf("n = %d and k = %d => kthBit:%c", n, k, findKthBit(n, k))
	fmt.Println()

	n = 5
	k = 16
	fmt.Printf("n = %d and k = %d => kthBit:%c", n, k, findKthBit(n, k))
	fmt.Println()

	n = 20
	k = 1048576 - 1
	fmt.Printf("n = %d and k = %d => kthBit:%c", n, k, findKthBit(n, k))
	fmt.Println()
}

func findKthBit(n int, k int) byte {
	if n == 1 || k == 1 {
		return '0'
	}
	if 2 <= k && k <= 4 {
		return '1'
	}
	sLength := int(math.Pow(2, float64(n)) - 1)
	mirrorIndex := sLength/2 + 1
	if mirrorIndex < k {
		// k is located on the right side of mirrorIndex
		reflectedIndex := sLength - k + 1
		n = n - 1
		return invertBit(findKthBit(n, reflectedIndex))
	}
	n = n - 1
	return findKthBit(n, k)
}

func invertBit(bit byte) byte {
	if bit == '0' {
		return '1'
	}
	return '0'
}
