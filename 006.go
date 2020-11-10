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
				getSquaredDiff(Input)
			}
		}
	}

}

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10  without any remainder.
What is the smallest positive number that is evenly divisible(divisible with no remainder) by all of the numbers from 1 to N ?
*/
func getSquaredDiff(N int64) {

	var sumOfN int64 = int64(math.Pow(float64(N*(N+1)/2), 2))
	var sumOfSquaredN int64 = N * (N + 1) * (2*N + 1) / 6
	fmt.Println(sumOfN - sumOfSquaredN)
}
