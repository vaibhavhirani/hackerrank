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
    * Complete the 'countApplesAndOranges' function below.
    *
    * The function accepts following parameters:
    *  1. INTEGER s
    *  2. INTEGER t
    *  3. INTEGER a
    *  4. INTEGER b
    *  5. INTEGER_ARRAY apples
    *  6. INTEGER_ARRAY oranges
    */
    /*
    s: integer, starting point of Sam's house location.
    t: integer, ending location of Sam's house location.
    a: integer, location of the Apple tree.
    b: integer, location of the Orange tree.
    apples: integer array, distances at which each apple falls from the tree.
    oranges: integer array, distances at which each orange falls from the tree.
    */

    func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
        // Write your code here
        noOfApples := 0
        noOfOranges := 0

        for i := 0; i < len(apples); i++ {
            apple :=  a + apples[i]
            if (apple >= s) && (apple <= t) {
                noOfApples = noOfApples + 1
            }  
        }

        for i := 0; i < len(oranges); i++ {
            orange :=  b + oranges[i]
            if (orange >= s) && (orange <= t) {
                noOfOranges = noOfOranges + 1
            }  
        }
        fmt.Println(noOfApples)
        fmt.Println(noOfOranges)
    }

    func main() {
        reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        sTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        s := int32(sTemp)

        tTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        t := int32(tTemp)

        secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        aTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
        checkError(err)
        a := int32(aTemp)

        bTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
        checkError(err)
        b := int32(bTemp)

        thirdMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        mTemp, err := strconv.ParseInt(thirdMultipleInput[0], 10, 64)
        checkError(err)
        m := int32(mTemp)

        nTemp, err := strconv.ParseInt(thirdMultipleInput[1], 10, 64)
        checkError(err)
        n := int32(nTemp)

        applesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var apples []int32

        for i := 0; i < int(m); i++ {
            applesItemTemp, err := strconv.ParseInt(applesTemp[i], 10, 64)
            checkError(err)
            applesItem := int32(applesItemTemp)
            apples = append(apples, applesItem)
        }

        orangesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var oranges []int32

        for i := 0; i < int(n); i++ {
            orangesItemTemp, err := strconv.ParseInt(orangesTemp[i], 10, 64)
            checkError(err)
            orangesItem := int32(orangesItemTemp)
            oranges = append(oranges, orangesItem)
        }

        countApplesAndOranges(s, t, a, b, apples, oranges)
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
