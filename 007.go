package main

import (
	"bufio"
	"fmt"
	//"math"
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
				getPrimeNumberAtArr(Input)
			}
		}
	}

}

/*
By listing the first six prime numbers: 2,3,5,7,11 and 13, we can see that the 12 prime is 6th.
What is the  nth prime number
*/
func getPrimeNumberAtArr(position int64) {
	var primeArray = []int64{2}
	var prime int64= 3
	var count int64= 0
	var i int= 0
	var primeFlag = false
	if (position == 1) {
		fmt.Println(2)
	} else if (position == 2){
		fmt.Println(3)
	} else {
		optimisedSlice := primeArray
		for count <= position-2 {
			primeFlag = false
			if (len(primeArray) >= 2) {
				divisionPoint := 0
				for i = 0; i < len(primeArray); i++ {
					if (primeArray[i] > prime/2) {
						divisionPoint = i
						break
					}
				}
				optimisedSlice = primeArray[:divisionPoint]
				// fmt.Println("first half", optimisedSlice)
			}
			for i = 0; i < len(optimisedSlice); i++ {
				if (prime%optimisedSlice[i] == 0) {
					//Number is divisible, therefore not a prime number
					primeFlag = false
					prime = prime + 2
					break
				} else {
					// fmt.Printf("%d / %d \n",prime, primeArray)
					primeFlag = true
				}
			}
			if (primeFlag == true) {
				primeArray = append(primeArray, prime)
				prime = prime + 2
				count = count + 1
			}
		}
	
		fmt.Println(primeArray[len(primeArray)-1])
	}
}
/*
2 3 5 7 11 13 17 19 23 29 31 37 
1 2 3 4 5  6  7  8  9  10 11 12

Algorithm - 

*/

