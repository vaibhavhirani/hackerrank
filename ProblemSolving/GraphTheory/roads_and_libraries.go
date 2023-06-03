package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'roadsAndLibraries' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER c_lib
 *  3. INinitTEGER c_road
 *  4. 2D_INTEGER_ARRAY cities
 */
var (
	library = make(map[int32]bool)
	roads   = make(map[int32][]int32)
)

func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	// Write your code here
	emptySlice := make([]int32, 0, 0)
	mapOfCity := make(map[int32][]int32)
	visit := make([]int32, n, n*int32(len(cities)))

	for i := 1; i <= int(n); i++ {
		mapOfCity[int32(i)] = emptySlice
		visit[i-1] = int32(i)
	}
	for i := 0; i < len(cities); i++ {
		pathTo := make([]int32, 0, n*int32(len(cities)))
		pathTo1 := make([]int32, 0, n*int32(len(cities)))
		pathTo = append(mapOfCity[cities[i][0]], cities[i][1])
		mapOfCity[int32(cities[i][0])] = pathTo
		pathTo1 = append(mapOfCity[cities[i][1]], cities[i][0])
		mapOfCity[int32(cities[i][1])] = pathTo1
	}

	// fmt.Printf("Cities %v \n", cities)
	// fmt.Printf("C -lib %v \n", c_lib)
	// fmt.Printf("C- Road %v \n", c_road)
	// fmt.Printf("Cities connected in  manner %v\n", mapOfCity)
	var cost int64
	var minimumCost int64
	for i := 0; i < len(visit); i++ {
		library = make(map[int32]bool)
		roads = make(map[int32][]int32)
		cost = calculateCost(visit, mapOfCity, c_lib, c_road)
		if i == 0 {
			minimumCost = cost
		}
		firstElement := visit[0]
		ithElement := visit[i]
		visit[0] = ithElement
		visit[i] = firstElement
		// fmt.Printf("%v > %v \n", minimumCost, cost)
		if minimumCost >= cost {
			minimumCost = cost
			// fmt.Printf("Minimun %v  \n", minimumCost)
		}
	}
	return minimumCost
}

func calculateCost(visit []int32, mapOfCityCircular map[int32][]int32, c_lib int32, c_road int32) int64 {
	var cost int32
	for i := 0; i < len(visit); i++ {
		l := int32(visit[i])
		k := mapOfCityCircular[l]
		// fmt.Printf("Checking for City %v\n", l)
		// fmt.Printf("%v can be connected to  %v\n", l, k)
		// fmt.Printf("%v is connected to  %v\n", l, mapOfCityCircular[l])
		// fmt.Printf("Initial Cost %v\n", cost)

		if hasLibrary(l) {
			continue
		} else {
			isConnectedToLib := false
			for _, j := range k {
				if hasLibrary(j) {
					if c_road <= c_lib {
						// fmt.Printf("2. Adding road  %v ---- %v \n", l, j)
						addRoadTo(l, j)
						cost = cost + c_road
						// fmt.Printf("3. Updating cost to %v after road addition \n\n", cost)
						isConnectedToLib = true
					} else {
						// fmt.Printf("4. Adding lib at  %v \n", l)
						addLibrary(l)
						cost = cost + c_lib
						// fmt.Printf("5. Updating cost to %v after lib addition \n\n", cost)
						isConnectedToLib = true
					}
					break
				}
			}
			if isConnectedToLib == false {
				// fmt.Printf("6. Adding lib at  %v \n", l)
				addLibrary(l)
				cost = cost + c_lib
				// fmt.Printf("7. Updating cost to %v after lib addition \n\n", cost)
			}
		}

	}
	// fmt.Printf("COST when starting with %v : %v \n\n\n\n", mapOfCityCircular, cost)
	return int64(cost)
}

func addRoadTo(a int32, b int32) {
	roads[a] = append(roads[a], b)
}

func addLibrary(a int32) {
	library[a] = true
}
func hasRoadTo(a int32, b int32) bool {
	d := int(b)
	for j := range roads[a] {
		if d == j {
			return true
		}
	}
	return false
}

func hasLibrary(a int32) bool {
	return library[a]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		c_libTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
