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
	p := intValues[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(intValues); i++ {
		if intValues[i] < p {
			left = append(left, intValues[i])
		} else {
			right = append(right, intValues[i])
		}
	}
	for _, n := range(left) {
		fmt.Print(n, " ")
	}
	fmt.Print(p, " ")
	for _, n := range(right) {
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
