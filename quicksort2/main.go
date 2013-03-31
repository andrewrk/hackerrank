package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"io"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	arrayLine, err := reader.ReadString('\n')
	if err != io.EOF && err != nil { panic(err) }
	intValues := toIntSlice(strings.Fields(arrayLine))
	qSort(intValues)
}

func qSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	p := slice[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(slice); i++ {
		if slice[i] < p {
			left = append(left, slice[i])
		} else {
			right = append(right, slice[i])
		}
	}
	left = qSort(left)
	right = qSort(right)
	result := make([]int, 0, len(left) + len(right) + 1)
	for _, n := range(left) {
		result = append(result, n)
	}
	result = append(result, p)
	for _, n := range(right) {
		result = append(result, n)
	}
	printNumbers(result)
	return result
}

func printNumbers(slice []int) {
	for _, n := range(slice) {
		fmt.Print(n, " ")
	}
	fmt.Print("\n")
}

func toIntSlice(slice []string) []int {
	intValues := make([]int, len(slice))
	for i, numStr := range(slice) {
		result, err := strconv.ParseInt(strings.TrimSpace(numStr), 10, 32)
		if err != nil { panic(err) }
		intValues[i] = int(result)
	}
	return intValues
}
