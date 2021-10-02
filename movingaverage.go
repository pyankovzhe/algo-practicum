package main

// Input example
// 1 2 3 4 5 6 7
// 4
// Output example
// 2.5 3.5 4.5 5.5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func movingAverage(array []int, windowSize int) []float64 {
	result := make([]float64, len(array)-(windowSize-1))
	var currentSum float64

	for i := 0; i < windowSize; i++ {
		currentSum += float64(array[i])
	}

	result[0] = currentSum / float64(windowSize)

	for i := 0; i < len(array)-windowSize; i++ {
		currentSum -= float64(array[i])
		currentSum += float64(array[i+windowSize])

		result[i+1] = float64(currentSum) / float64(windowSize)
	}

	return result
}

func main() {
	scanner := makeScanner()
	array := readArray(scanner)
	windowSize := readInt(scanner)
	printArray(movingAverage(array, windowSize))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []float64) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.FormatFloat(arr[i], 'f', 8, 64))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
