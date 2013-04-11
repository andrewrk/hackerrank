package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
)

type Pos struct {
	x int
	y int
}

func (p * Pos) Print() {
	// x and y are swapped because that's how the problem definition says to do it
	fmt.Println(p.y, p.x)
}

type World [][]uint8

var neighbors = [4]Pos{
	Pos{-1, 0},
	Pos{ 1, 0},
	Pos{ 0,-1},
	Pos{ 0, 1},
}

func readWorld(stream io.Reader) World {
	reader := bufio.NewReader(stream)
	// trash the first line
	_, err := reader.ReadString('\n')
	if err != nil { panic(err) }

	eof := false
	cells := make([][]uint8, 0)
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
		row := make([]uint8, len(line))
		cells = append(cells, row)
		for x, char := range line {
			row[x] = uint8(char)
		}
	}
	return cells
}

func (w World) ComputeBestMove() Pos {
	// identify groups
	type group []Pos
	groups := make([]group, 0)
	closedCells := NewWorld(len(w[0]), len(w))
	for y, row := range(w) {
		for x, c := range(row) {
			if c == '-' || closedCells[y][x] == 1 {
				continue
			}
			g := group{}
			openNodes := []Pos{ Pos{x, y} }
			for len(openNodes) > 0 {
				// pop openNodes
				n := openNodes[len(openNodes) - 1]
				openNodes = openNodes[:len(openNodes) - 1]
				nodeChar := w[n.y][n.x];
				if nodeChar == c {
					closedCells[n.y][n.x] = 1
					g = append(g, n)
					// add neighbors to open nodes
					for _, p := range(neighbors) {
						x2 := n.x + p.x
						y2 := n.y + p.y
						if x2 >= 0 && x2 < len(row) && y2 >=0 && y2 < len(w) && closedCells[y2][x2] == 0 {
							openNodes = append(openNodes, Pos{x2, y2})
						}
					}
				}
			}
			groups = append(groups, g)
		}
	}
	fmt.Println(groups)
	// TODO
	return Pos{}
}

func (w World) Step(move Pos) World {
	// TODO
	return w
}

func (w World) Clone() World {
	nw := NewWorld(len(w[0]), len(w))
	for y, row := range(w) {
		for x, c := range(row) {
			nw[y][x] = c
		}
	}
	return nw
}

func NewWorld(width, height int) World {
	w := make([][]uint8, height)
	for y := range(w) {
		w[y] = make([]uint8, width)
	}
	return w
}

func main () {
	w := readWorld(os.Stdin)
	move := w.ComputeBestMove()
	move.Print()
}

