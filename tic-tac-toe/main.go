package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

const (
	EMPTY int = iota
	O
	X
)
var otherChar = [3]int{EMPTY, X, O}


type Pos struct {
	x int
	y int
}

func (p * Pos) print() {
	// x and y are swapped because that's
	// how the problem definition says to do it
	fmt.Println(p.y, p.x)
}

type Board struct {
	myChar int
	cells [3][3]int
}

func (b * Board) print() {
	fmt.Println(b.myChar)
	fmt.Println(b.cells)
}

func (b * Board) winner() int {
	c00 := b.cells[0][0]
	c01 := b.cells[0][1]
	c02 := b.cells[0][2]

	c10 := b.cells[1][0]
	c11 := b.cells[1][1]
	c12 := b.cells[1][2]

	c20 := b.cells[2][0]
	c21 := b.cells[2][1]
	c22 := b.cells[2][2]

	// horizontal
	if c00 == c01 && c00 == c02 { return c00 }
	if c10 == c11 && c10 == c12 { return c10 }
	if c20 == c21 && c20 == c22 { return c20 }
	// vertical
	if c00 == c10 && c00 == c20 { return c00 }
	if c01 == c11 && c01 == c21 { return c01 }
	if c02 == c12 && c02 == c22 { return c02 }
	// diagonal
	if c00 == c11 && c00 == c22 { return c00 }
	if c02 == c11 && c02 == c20 { return c02 }

	return EMPTY
}

func parseMark(b byte) int {
	switch b {
	case '_': return EMPTY
	case 'X': return X
	case 'O': return O
	}
	panic("cannot parse mark")
}

func readBoard(stream io.Reader) Board {
	reader := bufio.NewReader(stream)
	bytes, err := reader.ReadBytes(0)
	if err != io.EOF { panic(err) }
	var board Board
	board.myChar = parseMark(bytes[0])
	board.cells[0][0] = parseMark(bytes[2])
	board.cells[0][1] = parseMark(bytes[3])
	board.cells[0][2] = parseMark(bytes[4])

	board.cells[1][0] = parseMark(bytes[6])
	board.cells[1][1] = parseMark(bytes[7])
	board.cells[1][2] = parseMark(bytes[8])

	board.cells[2][0] = parseMark(bytes[10])
	board.cells[2][1] = parseMark(bytes[11])
	board.cells[2][2] = parseMark(bytes[12])
	return board
}

func (b Board) nextMove() (bestScore int, scoreGrid [3][3]int, bestMove Pos) {
	switch b.winner() {
	case b.myChar:
		bestScore = 1
		return
	case otherChar[b.myChar]:
		bestScore = -1
		return
	}
	foundEmpty := false
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if b.cells[y][x] == EMPTY {
				// evaluate the goodness of this move
				nextBoard := b
				nextBoard.myChar = otherChar[b.myChar]
				nextBoard.cells[y][x] = b.myChar
				opponentScore, _, _ := nextBoard.nextMove()
				scoreGrid[y][x] = -opponentScore
				if !foundEmpty || -opponentScore > bestScore {
					foundEmpty = true
					bestScore = -opponentScore
					bestMove = Pos{x, y}
				}
			}
		}
	}
	return
}

func main() {
	board := readBoard(os.Stdin)
	_, _, move := board.nextMove()
	move.print()
}
