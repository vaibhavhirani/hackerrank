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
				getLargestPrimeFactor(Input)
			}
		}
	}

}

func getLargestPrimeFactor(num int64) {
	var i int64 = 0
	var divisor int64

	_, frac := math.Modf(math.Sqrt(float64(num)))
	if frac == 0 {
		num = int64(math.Sqrt(float64(num)))
	}

	for num%2 == 0 {
		divisor = num
		num = num / 2
	}

	for i = 3; i <= int64(math.Sqrt(float64(num))); i = i + 2 {
		_, frac = math.Modf(math.Sqrt(float64(num)))
		if frac == 0 {
			num = int64(math.Sqrt(float64(num)))
		}
		for num%i == 0 {
			divisor = num
			num = num / i
		}
	}

	if num == 1 {
		fmt.Println(divisor)
	} else {
		fmt.Println(num)
	}
}
