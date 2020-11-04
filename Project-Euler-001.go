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
			if Input-1000000000 > 0 {
				os.Exit(0)
			} else {
				method(Input)
			}
		}
	}

}

func method(N int64) {
	output := ap(N, 5) + ap(N, 3) - ap(N, 15)
	fmt.Println(output)
}

func ap(N int64, d int64) int64 {
	var n int64
	if N%d == 0 {
		n = int64(float64(N) / float64(d))
		n = n - 1

	} else {
		n = int64(math.Abs(float64(N) / float64(d)))
	}
	result := (2 * d) + (n-1)*d
	result = n * result / 2
	return result
}
