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
	sum := 0
	for x := 1; x < len(intValues); x++ {
		tmp := intValues[x]
		var i int
		for i = x - 1; i >= 0; i-- {
			if tmp < intValues[i] {
				sum += 1
				intValues[i + 1] = intValues[i]
			} else {
				break
			}
		}
		intValues[i + 1] = tmp
	}
	fmt.Println(sum)
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
