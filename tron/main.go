package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"io"
	"strings"
)

const (
	EMPTY int = iota
	WALL
	RED
	GREEN
)

var otherChar = [...]int{EMPTY, WALL, GREEN, RED}

type Pos struct {
	x int
	y int
}

type World struct {
	myPos Pos
	myChar int
	oPos Pos
	cells [][]int
}

func (w * World) print() {
	fmt.Println(w.myChar, w.myPos)
	fmt.Println(otherChar[w.myChar], w.oPos)
	for _, row := range(w.cells) {
		for _, c := range(row) {
			fmt.Print(c)
		}
		fmt.Print("\n")
	}
}

func parsePos(y, x string) Pos {
	// x and y swapped 'cause that's how hackerrank rolls
	x64, err := strconv.ParseInt(x, 10, 32)
	if err != nil { panic(err) }
	y64, err := strconv.ParseInt(y, 10, 32)
	if err != nil { panic(err) }
	return Pos{int(x64), int(y64)}
}

func parseChar(b uint8) int {
	switch b {
	case 'r': return RED
	case 'g': return GREEN
	case '#': return WALL
	case '-': return EMPTY
	}
	panic("bad parse char")
}

func readWorld(stream io.Reader) * World {
	reader := bufio.NewReader(stream)
	myCharLine, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	positionsLine, err := reader.ReadString('\n')
	posList := strings.Fields(positionsLine)
	if err != nil { panic(err) }
	w := new(World)
	w.myChar = parseChar(myCharLine[0])
	w.myPos = parsePos(posList[0], posList[1])
	w.oPos = parsePos(posList[2], posList[3])
	w.cells = make([][]int, 0)
	eof := false
	for !eof {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			eof = true
			if len(line) == 0 {
				break
			}
		} else if err != nil {
			panic(err)
		}
		row := make([]int, len(line))
		w.cells = append(w.cells, row)
		for x, char := range line {
			row[x] = parseChar(uint8(char))
		}
	}
	return w
}

func main() {
	w := readWorld(os.Stdin)
	if w.myPos.y == 7 {
		if w.myChar == RED && w.cells[w.myPos.y][w.myPos.x + 1] == EMPTY {
			fmt.Println("RIGHT")
			return
		} else if w.myChar == GREEN && w.cells[w.myPos.y][w.myPos.x - 1] == EMPTY {
			fmt.Println("LEFT")
			return
		}
	}

}
