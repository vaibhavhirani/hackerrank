package main

import (
	"bufio"
	"fmt"
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
			//Input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			Input, _ := strconv.Atoi(scanner.Text())
			if Input-40000000000000000 > 0 {
				os.Exit(0)
			} else {
				genFibonacciSeries(Input)
			}
		}
	}

}

func genFibonacciSeries(term int) {
	firstTerm := 1
	secondTerm := 2
	//while loop
	thirdTerm := firstTerm + secondTerm
	sum := 2
	for thirdTerm < term {
		firstTerm = secondTerm
		secondTerm = thirdTerm
		if thirdTerm%2 == 0 {
			sum = sum + thirdTerm
		}
		thirdTerm = firstTerm + secondTerm
	}
	fmt.Println(sum)
}
