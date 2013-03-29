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

func (w * world) dirtyCells() []pos {
	cells := make([]pos, 0)
	var it pos
	for it.y = 0; it.y < len(w.cells); it.y++ {
		for it.x = 0; it.x < len(w.cells[it.y]); it.x++ {
			if w.cells[it.y][it.x] == 'd' {
				cells = append(cells, it)
			}
		}
	}

	return cells
}

func (w * world) clone() * world {
	newStruct := *w
	posClone := *newStruct.botPos
	newStruct.botPos = &posClone

	newStruct.cells = make([][]rune, len(w.cells))
	for y, row := range(w.cells) {
		newStruct.cells[y] = make([]rune, len(row))
		for x, char := range(row) {
			newStruct.cells[y][x] = char
		}
	}

	return &newStruct;
}

func (w * world) hideInvisible() {
	for y := 0; y < w.botPos.y - 1; y++ {
		w.hideInvisibleRow(y)
	}
	for y := w.botPos.y + 2; y < len(w.cells); y++ {
		w.hideInvisibleRow(y)
	}
}

func (w * world) hideInvisibleRow(y int) {
	for x := 0; x < w.botPos.x - 1; x++ {
		w.hide(x, y)
	}
	for x := w.botPos.x + 2; x < len(w.cells[0]); x++ {
		w.hide(x, y)
	}
}

func (w * world) hide(x, y int) {
	w.cells[y][x] = 'o'
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

const (
	MOVE_LEFT = iota
	MOVE_RIGHT
	MOVE_UP
	MOVE_DOWN
	MOVE_CLEAN
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

func computeCost(moveMap[][]rune, w * world) int {
	sum := 0
	realSimWorld := w.clone() // the real data
	for len(realSimWorld.dirtyCells()) > 0 {
		simWorld := realSimWorld.clone()
		simWorld.hideInvisible()
		switch computeNextMove(moveMap, simWorld) {
		case MOVE_CLEAN:
			realSimWorld.cells[realSimWorld.botPos.y][realSimWorld.botPos.x] = '-'
		case MOVE_LEFT:
			realSimWorld.botPos.x -= 1
		case MOVE_RIGHT:
			realSimWorld.botPos.x += 1
		case MOVE_UP:
			realSimWorld.botPos.y -= 1
		case MOVE_DOWN:
			realSimWorld.botPos.y += 1
		}
		sum += 1
	}
	return sum
}

func computeNextMove(moveMap[][]rune, w * world) int {
	// always clean if we're in a dirty cell
	if w.cells[w.botPos.y][w.botPos.x] == 'd' {
		return MOVE_CLEAN;
	}
	// collect the known dirty spots
	dirtyCells := w.dirtyCells()

	if len(dirtyCells) == 0 {
		// no dirty spot visible. follow a prescribed track to hit all the cells
		switch moveMap[w.botPos.y][w.botPos.x] {
		case '>': return MOVE_RIGHT
		case '<': return MOVE_LEFT
		case '^': return MOVE_UP
		case 'v': return MOVE_DOWN
		}
		panic("unknown movemap value")
	}

	// figure out which dirty cell to go after
	var lowestCostCell * pos = nil
	var lowestCost int
	for _, dirtyCell := range(dirtyCells) {
		cost := rectilinearDist(&dirtyCell, w.botPos)
		// create a world in which we have moved to
		// and cleaned that cell
		simWorld := w.clone()
		simWorld.botPos = &dirtyCell
		simWorld.cells[dirtyCell.y][dirtyCell.x] = '-'
		cost += computeCost(moveMap, simWorld)
		if (lowestCostCell == nil || cost < lowestCost) {
			lowestCostCell = &dirtyCell
			lowestCost = cost
		}
	}
	if w.botPos.x < lowestCostCell.x {
		return MOVE_RIGHT
	} else if w.botPos.x > lowestCostCell.x {
		return MOVE_LEFT
	} else if w.botPos.y < lowestCostCell.y {
		return MOVE_DOWN
	} else if w.botPos.y > lowestCostCell.y {
		return MOVE_UP
	}
	panic("shouldn't get here")
}

func main() {
	w, err := readWorld(os.Stdin)
	if err != nil { panic(err) }
	moveMap := buildMovementMap(w.cells)
	switch computeNextMove(moveMap, w) {
	case MOVE_LEFT: fmt.Println("LEFT")
	case MOVE_RIGHT: fmt.Println("RIGHT")
	case MOVE_UP: fmt.Println("UP")
	case MOVE_DOWN: fmt.Println("DOWN")
	case MOVE_CLEAN: fmt.Println("CLEAN")
	}

}

