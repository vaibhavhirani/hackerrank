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
			Input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			// Input, _ := strconv.Atoi(scanner.Text())
			if Input-1000000000000 > 0 {
				os.Exit(0)
			} else {
				getSmallestMul(Input)
			}
		}
	}

}

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10  without any remainder.
What is the smallest positive number that is evenly divisible(divisible with no remainder) by all of the numbers from 1 to N ?
*/
func getSmallestMul(N int64) {
	var i, j, smallest, count int64 = 0, 1, 0, 0

	for N != count {
		// fmt.Printf("Smallest %v \n", smallest)
		smallest = N * j
		count = 0
		for i = N; i >= 1; i-- {
			// fmt.Printf("Dividing %v \n", smallest)
			if smallest%i == 0 {
				// fmt.Printf("\t Dividing %v / %v\n", smallest, i)
				count = count + 1
				// fmt.Printf("\t Count %v\n", count)
			} else {
				count = 0
				break
			}
			// fmt.Printf("\t Count %v\n", count)
		}
		j = j + 1
	}
	fmt.Println(smallest)
}
