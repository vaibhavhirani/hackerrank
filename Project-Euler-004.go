package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	if N-100000 > 0 {
		os.Exit(0)
	} else {
		for i := 0; i < N; i++ {
			scanner.Scan()
			Input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			// Input, _ := strconv.Atoi(scanner.Text())
			if Input-1000000000000 > 0 {
				os.Exit(0)
			} else {
				getLargestPalindrome(Input)
			}
		}
	}

}

// checks whether generated number is palindrome
func getLargestPalindrome(num int64) {

	var i, j int64
	var product int64
	var prevPalindrome int64
	var reverseOfProduct string
	var reverseOfProductInt int64
	var largestPalindrome int64 = 101101
	var k int64
	var count int64
	var realJ int64
	var limit int64 = 500
	var oldLimit, temp int64 = 0, 0

	for i = 121; i < 1000; i++ {
		if int64(math.Abs(float64(100000/i)))-i < 0 {
			realJ = i
		} else {
			realJ = int64(math.Abs(float64(100000 / i)))
		}
		if (i+realJ)%2 != 0 {
			realJ = realJ - 1
		}
		fmt.Println(realJ)
		for j = realJ; j < 1000; j = j + 2 {
			count = count + 1
			if i*j >= num {
				break
			}
			product = i * j
			prevPalindrome = product
			// fmt.Printf("Product of  %v * %v is %v \n", i, j, product)
			// String approach
			reverseOfProduct = ""
			for k = 5; k >= 0; k-- {
				digit := prevPalindrome % 10
				prevPalindrome = prevPalindrome / 10
				digitStr := strconv.FormatInt(digit, 10)
				reverseOfProduct = reverseOfProduct + digitStr
				// reverseOfProduct = reverseOfProduct + digit*(int64(math.Pow(float64(ten), float64(k))))
			}
			// temp = largestPalindrome //last palindrome
			reverseOfProductInt, _ = strconv.ParseInt(reverseOfProduct, 10, 64)
			if product == reverseOfProductInt {
				// fmt.Printf("Product of  %v * %v Reverse of %v is %v \n", i, j, product, reverseOfProductInt)
				if largestPalindrome < product {
					largestPalindrome = product
				}
			}
		}
		// fmt.Println()
	}
	fmt.Println(largestPalindrome)
	// fmt.Printf("Count %v \n", count)
}
