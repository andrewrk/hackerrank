package main

import (
	"fmt"
	"bufio"
	"time"
	"os"
	"io"
	"strings"
)

const (
	DEAD int = iota
	BLACK
	WHITE
)

var otherColor = [3]int{DEAD, WHITE, BLACK}

type cell struct {
	value int
	neighborCount [3]int
}

type pos struct {
	x int
	y int
}

func (p * pos) print() {
	// x and y are swapped because that's how the problem definition says to do it
	fmt.Println(p.y, p.x)
}

func (c * cell) alive() bool {
	return c.value != DEAD
}

func (c * cell) useful() bool {
	return c.neighborCount[WHITE] + c.neighborCount[BLACK] > 0
}

type ScoreGrid [][]int

type world struct {
	myColor int
	cells [][]cell
}

func parseColor (color uint8) int {
	switch color {
	case 'w': return WHITE
	case 'b': return BLACK
	case '-': return DEAD
	}
	panic("bad color")
}

func (w * world) inBounds(x, y int) bool {
	if y < 0 || x < 0 { return false }
	if y >= len(w.cells) || x >= len(w.cells[0]) { return false }
	return true
}

func (w * world) step() * world {
	nw := newWorld(len(w.cells[0]), len(w.cells), w.myColor)
	for y, row := range(w.cells) {
		for x, c := range(row) {
			total := c.neighborCount[WHITE] + c.neighborCount[BLACK]
			switch {
			case c.value == DEAD && total == 3:
				nw.cells[y][x].value = WHITE
				if c.neighborCount[BLACK] > c.neighborCount[WHITE] {
					nw.cells[y][x].value = BLACK
				}
			case (c.value == WHITE || c.value == BLACK) && total != 2 && total != 3:
				nw.cells[y][x].value = DEAD
			default:
				nw.cells[y][x].value = c.value

			}
		}
	}
	nw.computeNeighborCounts()
	return nw
}

func (w * world) score() [3]int {
	scores := [3]int{0, 0, 0}
	for _, row := range(w.cells) {
		for _, c := range(row) {
			scores[c.value] += 1
		}
	}
	return scores
}

func (w * world) print() {
	fmt.Println(w.myColor)
	for _, row := range(w.cells) {
		for _, c := range(row) {
			fmt.Print(c.value)
			//fmt.Print(c.neighborCount[WHITE] + c.neighborCount[BLACK])
		}
		fmt.Print("\n")
	}
}

func (w * world) clone()  * world {
	nw := newWorld(len(w.cells[0]), len(w.cells), w.myColor)
	for y, row := range(w.cells) {
		for x, c := range(row) {
			nw.cells[y][x] = c
		}
	}
	return nw
}

func NewScoreGrid(w, h int) ScoreGrid {
	sg := make([][]int, h)
	for y := range(sg) {
		sg[y] = make([]int, h)
	}
	return sg
}

func (sg ScoreGrid) print() {
	for _, row := range(sg) {
		for _, v := range(row) {
			fmt.Print(v, " ")
		}
		fmt.Print("\n")
	}
}

func (w * world) computeNeighborCounts() {
	// reset
	for _, row := range(w.cells) {
		for x := range(row) {
			row[x].neighborCount[DEAD] = 0
			row[x].neighborCount[BLACK] = 0
			row[x].neighborCount[WHITE] = 0
		}
	}
	// count
	for y, row := range(w.cells) {
		for x := range(row) {
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					nx := x + dx
					ny := y + dy
					if (dx != 0 || dy != 0) && w.inBounds(nx, ny) {
						w.cells[ny][nx].neighborCount[row[x].value] += 1
					}
				}
			}

		}
	}
}

func newWorld(width, height, myColor int) * world {
	w := new(world)
	w.myColor = myColor
	w.cells = make([][]cell, height)
	for y := range(w.cells) {
		w.cells[y] = make([]cell, width)
		for x := range(w.cells[y]) {
			w.cells[y][x].neighborCount[DEAD] = 8
		}
	}
	return w
}

func readWorld(stream io.Reader) (*world, error) {
	reader := bufio.NewReader(stream)
	myCharLine, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	w := new(world)
	w.myColor = parseColor(myCharLine[0])
	w.cells = make([][]cell, 0)
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
		row := make([]cell, len(line))
		w.cells = append(w.cells, row)
		for x, char := range line {
			row[x].value = parseColor(uint8(char))
		}
	}
	w.computeNeighborCounts()
	return w, nil
}

func main () {
	timeLimit := float64(5.0)
	startTime := time.Now()

	original, err := readWorld(os.Stdin)
	if err != nil { panic(err) }
	// try every possibility and check how it performs
	sg := NewScoreGrid(len(original.cells[0]), len(original.cells))
	first := true
	var best int
	var bestPos pos
	secondBestPos := pos{len(original.cells[0]) / 2, len(original.cells) / 2}

	for y, row := range(original.cells) {
		for x, c := range(row) {
			timePassed := time.Since(startTime)
			if timePassed.Seconds() >= timeLimit {
				break
			}
			if c.value == DEAD && c.useful() {
				// try placing here
				w := original.clone()
				w.cells[y][x].value = w.myColor
				// step forward in time
				for i := 0; i < 500; i++ {
					w = w.step()
				}
				scores := w.score()
				score := scores[w.myColor] - scores[otherColor[w.myColor]]
				sg[y][x] = score
				if first || score > best {
					first = false
					best = score
					bestPos = pos{x, y}
				}
			} else if c.value != DEAD && x + 1 < len(original.cells[0]) && original.cells[y][x + 1].value == DEAD {
				secondBestPos = pos{x + 1, y}
			}
		}
	}
	if first {
		// no good move found, use fallback
		secondBestPos.print()
	} else {
		bestPos.print()
		//fmt.Println("score:", best)
	}
	//sg.print()
}
