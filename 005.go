package main

import (
	"bufio"
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
