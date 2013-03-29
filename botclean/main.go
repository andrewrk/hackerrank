package main

import (
	"os"
	"bufio"
	"io"
	"strconv"
	"fmt"
	"strings"
)

type pos struct {
	x int
	y int
}

type world struct {
	botPos pos
	cells [][]bool
}

func parsePos(line string) (*pos, error) {
	slice := strings.Split(strings.TrimSpace(line), " ")
	// flip x and y
	result := new(pos)
	x64, err := strconv.ParseInt(slice[1], 10, 32)
	if err != nil { return nil, err }
	y64, err := strconv.ParseInt(slice[0], 10, 32)
	if err != nil { return nil, err }
	result.x = int(x64)
	result.y = int(y64)
	return result, nil
}

func readWorld(reader io.Reader) (*world, error) {
	reader := bufio.NewReader(stream)
	botPosLine, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	botPos, err := parsePos(botPosLine)
	if err != nil { panic(err) }
	grid := make([][]bool, 0, 100)
	eof := false
	for y := 0; !eof; y++ {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			eof = true
		} else if err != nil {
			panic(err)
		}
		row := make([]bool, len(line), len(line))
		grid = append(grid, row)
		for x, char := range line {
			if char == 'd' {
				row[x] = true
			}
		}
	}
}

func main() {
	w, err := readWorld(os.Stdin)
	if err != nil { panic(err) }
}
