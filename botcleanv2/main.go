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
	botPos * pos
	cells [][]rune
}

func parsePos(line string) (*pos, error) {
	slice := strings.Split(strings.TrimSpace(line), " ")
	// flip x and y
	x64, err := strconv.ParseInt(slice[1], 10, 32)
	if err != nil { return nil, err }
	y64, err := strconv.ParseInt(slice[0], 10, 32)
	if err != nil { return nil, err }
	return &pos{int(x64), int(y64)}, nil
}

func readWorld(stream io.Reader) (*world, error) {
	reader := bufio.NewReader(stream)
	botPosLine, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	w := new(world)
	w.botPos, err = parsePos(botPosLine)
	if err != nil { panic(err) }
	w.cells = make([][]rune, 0)
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
		row := make([]rune, len(line))
		w.cells = append(w.cells, row)
		for x, char := range line {
			row[x] = char
		}
	}
	return w, nil
}

type posMatch struct {
	p * pos
	distance int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func rectilinearDist(pos1 * pos, pos2 * pos) int {
	return abs(pos1.x - pos2.x) + abs(pos1.y - pos2.y)
}

/*
generate these:

(5x5)
vvvvv
>>>v<
^vvv<
^vvv<
^<<<<

(6x6)
vvvvvv
>>>>v<
^vvvv<
^vvvv<
^<<<<<
^<<<<<

(7x7)
vvvvvvv
>>>>>v<
^vvvvv<
^vvvvv<
^v<<<<<
^>>>>v<
^<<<<<<

(8x8)
vvvvvvvv
>>>>>>v<
^vvvvvv<
^vvvvvv<
^v<<<<<<
^vvvvvv<
^>>>>>v<
^<<<<<<<

(9x9)
vvvvvvvvv
>>>>>>>v<
^vvvvvvv<
^vvvvvvv<
^v<<<<<<<
^vvvvvvv<
^vvvvvvv<
^>>>>>>v<
^<<<<<<<<

(10x10)
vvvvvvvvvvv
>>>>>>>>>v<
^vvvvvvvvv<
^vvvvvvvvv<
^v<<<<<<<<<
^vvvvvvvvv<
^vvvvvvvvv<
^>>>>>>>>v<
^vvvvvvvvv<
^vvvvvvvvv<
^<<<<<<<<<<

*/

const (
	STATE_RIGHT = iota
	STATE_LEFT
)

func buildMovementMap(cells [][]rune) [][]rune {
	moveMap := make([][]rune, len(cells))
	// fill grid with 'v'
	for y, row := range(cells) {
		moveMap[y] = make([]rune, len(row))
		for x := range(row) {
			moveMap[y][x] = 'v'
		}
	}
	// fill right side with '<'
	right := len(cells[0]) - 1
	for y := range(cells) {
		moveMap[y][right] = '<'
	}
	// fill bottom row with '<'
	bottom := len(cells) - 1
	for x := range(cells[0]) {
		moveMap[bottom][x] = '<'
	}
	// left column
	for y := 2; y < len(cells); y++ {
		moveMap[y][0] = '^'
	}
	// entry on left column
	moveMap[1][0] = '>'
	// zig zag
	zz := pos{1, 1}
	state := STATE_RIGHT
	for {
		if state == STATE_RIGHT {
			moveMap[zz.y][zz.x] = '>'
			zz.x += 1
			if zz.x == right - 1 {
				state = STATE_LEFT
				zz.y += 3
				if zz.y >= bottom {
					break
				}
			}
		} else if state == STATE_LEFT {
			moveMap[zz.y][zz.x] = '<'
			zz.x -= 1
			if zz.x == 1 {
				if zz.y == bottom - 1 && zz.x == 1 {
					moveMap[zz.y][zz.x] = '<'
					break
				} else if zz.y == bottom - 2 && zz.x == 1 {
					for x := 1; x < right - 1; x++ {
						moveMap[bottom - 1][x] = '>'
					}
					moveMap[bottom - 1][right - 1] = 'v'
					break
				} else if zz.y == bottom - 3 && zz.x == 1 {
					for x := 1; x < right - 1; x++ {
						moveMap[bottom - 1][x] = '>'
					}
					moveMap[bottom - 1][right - 1] = 'v'
					break
				}
				state = STATE_RIGHT
				zz.y += 3
				if zz.y >= bottom {
					break
				}
			}
		}
	}
	return moveMap
}

func printMoveMap(moveMap[][]rune) {
	for _, row := range(moveMap) {
		for _, char := range(row) {
			fmt.Print(string(char))
		}
		fmt.Print("\n")
	}
}

func main() {
	w, err := readWorld(os.Stdin)
	if err != nil { panic(err) }
	// always clean if we're in a dirty cell
	if w.cells[w.botPos.y][w.botPos.x] == 'd' {
		fmt.Println("CLEAN")
		return;
	}
	// if we see a dirty spot, go clean it
	closest := new(posMatch)
	var it pos
	for it.y = 0; it.y < len(w.cells); it.y++ {
		for it.x = 0; it.x < len(w.cells[it.y]); it.x++ {
			if w.cells[it.y][it.x] == 'd' {
				// dirty
				dist := rectilinearDist(&it, w.botPos)
				if closest.p == nil || dist < closest.distance {
					var clone pos = it
					closest.p = &clone
					closest.distance = dist
				}
			}
		}
	}
	if closest.p != nil {
		if w.botPos.x < closest.p.x {
			fmt.Println("RIGHT")
		} else if w.botPos.x > closest.p.x {
			fmt.Println("LEFT")
		} else if w.botPos.y < closest.p.y {
			fmt.Println("DOWN")
		} else if w.botPos.y > closest.p.y {
			fmt.Println("UP")
		}
		return
	}
	// no dirty spot visible. follow a prescribed track to hit all the cells
	moveMap := buildMovementMap(w.cells)
	switch moveMap[w.botPos.y][w.botPos.x] {
	case '>': fmt.Println("RIGHT")
	case '<': fmt.Println("LEFT")
	case '^': fmt.Println("UP")
	case 'v': fmt.Println("DOWN")
	}
}

