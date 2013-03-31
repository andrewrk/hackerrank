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
	valuesLine, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	valueStrs := strings.Fields(valuesLine)
	delta, err := strconv.ParseInt(valueStrs[1], 10, 64)
	if err != nil { panic(err) }
	arrayLine, err := reader.ReadString('\n')
	if err != io.EOF && err != nil { panic(err) }
	intValues := toIntSlice(strings.Fields(arrayLine))
	table := map[int64] int {}
	sum := 0
	for _, n := range(intValues) {
		count, ok := table[n]
		if !ok {
			count = 0
		}
		sum += count
		left := n - delta
		prevLeft, ok := table[left]
		if ok {
			table[left] = prevLeft + 1
		} else {
			table[left] = 1
		}

		right := n + delta
		prevRight, ok := table[right]
		if ok {
			table[right] = prevRight + 1
		} else {
			table[right] = 1
		}
	}
	fmt.Println(sum)
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
