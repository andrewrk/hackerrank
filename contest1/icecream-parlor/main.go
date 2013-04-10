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
	testCountLine, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	testCount64, err := strconv.ParseInt(strings.TrimSpace(testCountLine), 10, 64)
	if err != nil { panic(err) }
	testCount := int(testCount64)
	for i := 0; i < testCount; i++ {
		test(reader)
	}
}

func test(reader * bufio.Reader) {
	costLine, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	cost, err := strconv.ParseInt(strings.TrimSpace(costLine), 10, 64)
	if err != nil { panic(err) }

	// trash the flavor count
	_, err = reader.ReadString('\n')
	if err != nil { panic(err) }

	arrayLine, err := reader.ReadString('\n')
	if err != io.EOF && err != nil { panic(err) }

	intValues := toIntSlice(strings.Fields(arrayLine))
	table := map[int64] int {}
	for index, n := range(intValues) {
		first, ok := table[n]
		if ok {
			fmt.Println(first + 1, index + 1)
			return
		}
		second := cost - n
		table[second] = index
	}
	fmt.Println("no solution")
}

func toIntSlice(slice []string) []int64 {
	intValues := make([]int64, len(slice))
	for i, numStr := range(slice) {
		result, err := strconv.ParseInt(strings.TrimSpace(numStr), 10, 64)
		if err != nil { panic(err) }
		intValues[i] = result
	}
	return intValues
}
