package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// with sorting
// func twoSum(array []int, targetSum int) []int {
// 	sort.Ints(array)

// 	i := 0
// 	j := len(array) - 1

// 	for i < j {
// 		sum := array[i] + array[j]
// 		if sum == targetSum {
// 			return []int{array[i], array[j]}
// 		} else if sum > targetSum {
// 			j--
// 		} else {
// 			i++
// 		}
// 	}

// 	return []int{}
// }

// with additional memory
func twoSum(array []int, targetSum int) []int {
	hash := make(map[int]int)

	for i, num := range array {
		diff := targetSum - num

		if val, ok := hash[diff]; ok {
			return []int{array[val], num}
		}

		hash[num] = i
	}

	return []int{}
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	array := readArray(scanner)
	targetSum := readInt(scanner)
	result := twoSum(array, targetSum)
	if len(result) == 0 {
		fmt.Println("None")
	} else {
		fmt.Print(result[0])
		fmt.Print(" ")
		fmt.Print(result[1])
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
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
