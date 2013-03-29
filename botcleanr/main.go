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
	cells [][]bool
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
	w.cells = make([][]bool, 0, 100)
	eof := false
	for !eof {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			eof = true
		} else if err != nil {
			panic(err)
		}
		row := make([]bool, len(line), len(line))
		w.cells = append(w.cells, row)
		for x, char := range line {
			if char == 'd' {
				row[x] = true
			}
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

func main() {
	w, err := readWorld(os.Stdin)
	if err != nil { panic(err) }
	// always clean if we're in a dirty cell
	if w.cells[w.botPos.y][w.botPos.x] {
		fmt.Println("CLEAN")
		return;
	}
	// simple strategy: go for the closest dirty spot
	closest := new(posMatch)
	var it pos
	for it.y = 0; it.y < len(w.cells); it.y++ {
		for it.x = 0; it.x < len(w.cells[it.y]); it.x++ {
			if w.cells[it.y][it.x] {
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
	if w.botPos.x < closest.p.x {
		fmt.Println("RIGHT")
	} else if w.botPos.x > closest.p.x {
		fmt.Println("LEFT")
	} else if w.botPos.y < closest.p.y {
		fmt.Println("DOWN")
	} else if w.botPos.y > closest.p.y {
		fmt.Println("UP")
	}
}

