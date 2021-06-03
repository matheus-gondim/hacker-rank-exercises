package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func removeIndex(s sort.IntSlice, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)    //vai de zero ao index
	return append(ret, s[index+1:]...) //vai do index ao final
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func miniMaxSum(arr []int32) {
	var intArray []int
	for _, num := range arr {
		intArray = append(intArray, int(num))
	}

	sort.Ints(intArray)

	removeLast := removeIndex(intArray, len(intArray)-1)
	removeFirst := removeIndex(intArray, 0)

	mini := sum(removeLast...)
	max := sum(removeFirst...)

	fmt.Println(mini, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
