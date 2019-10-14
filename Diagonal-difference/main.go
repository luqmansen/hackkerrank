package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {

	var ltr, rtl int32

	for i := 0; i < len(arr); i++ {

		ltr += arr[i][i]
		rtl += arr[i][len(arr)-1-i]

	}
	return int32(math.Abs(float64(ltr - rtl)))

}

func main() {

	arr := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	fmt.Println(diagonalDifference(arr))
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	//
	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//
	//defer stdout.Close()
	//
	//writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
	//
	//nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//n := int32(nTemp)
	//
	//var arr [][]int32
	//for i := 0; i < int(n); i++ {
	//	arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")
	//
	//	var arrRow []int32
	//	for _, arrRowItem := range arrRowTemp {
	//		arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
	//		checkError(err)
	//		arrItem := int32(arrItemTemp)
	//		arrRow = append(arrRow, arrItem)
	//	}
	//
	//	if len(arrRow) != int(n) {
	//		panic("Bad input")
	//	}
	//
	//	arr = append(arr, arrRow)
	//}
	//
	//result := diagonalDifference(arr)
	//
	//fmt.Fprintf(writer, "%d\n", result)
	//
	//writer.Flush()
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
