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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	isPM := strings.Contains(s, "PM")

	strClean := strings.ReplaceAll(s, "PM", "")
	strClean = strings.ReplaceAll(strClean, "AM", "")

	strSlice := strings.Split(strClean, ":")

	hour, _ := strconv.Atoi(strSlice[0])

	if isPM && hour != 12 {
		strSlice[0] = strconv.Itoa(hour + 12)
	}

	if !isPM && hour == 12 {
		strSlice[0] = "00"
	}

	finalStr := strings.Join(strSlice, ":")
	return finalStr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
