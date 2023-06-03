package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        k := int32(kTemp)

        num := readLine(reader)

        getGreatestProduct(n, k, num)
    }
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

func getGreatestProduct(n int32, k int32, line string) int32 {
    //check for even or odd length
    isItEven := n%2
    lenOfLine :=int(n/2)
    if isItEven == 1 {
        lenOfLine = lenOfLine + 1
    }
    product := 0
    // line is n long string 
    for i:=0; i < lenOfLine ; i++ {
         
         //calculates product here
         tempProduct := 1
         
        // divide it in k length number 
         number := string(line[i:i+int(k)])
         
        // find product of each number
         for j := 0; j < len(number); j++ {
            convertedNumber, _ := strconv.Atoi(string(number[j])) 
            tempProduct = tempProduct*convertedNumber
         }
         
        // return the greatest
        if tempProduct > product {
            product = tempProduct
        }
    }

    fmt.Println(n,k,line,product)
    fmt.Println(product)

    return n
}

