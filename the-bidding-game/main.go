package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"io"
	"strings"
)

type Game struct {
	cash [2]int
	bottlePos int
	drawAdvantage int
}

var otherPlayer = [2]int{1, 0}
var playerMoveDelta = [2]int{-1, 1}

func NewGame() Game {
	return Game{[2]int{100, 100}, 5, 0}
}

func (g * Game) next(bids [2]int) {
	var winner int
	if bids[0] > bids[1] {
		winner = 0
	} else if bids[1] > bids[0] {
		winner = 1
	} else {
		winner = g.drawAdvantage
		g.drawAdvantage = otherPlayer[winner]
	}
	g.cash[winner] -= bids[winner]
	g.bottlePos += playerMoveDelta[winner]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	playerLine, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	scotchPosLine, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	p1BidsLine, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	p2BidsLine, err := reader.ReadString('\n')
	if err != io.EOF && err != nil { panic(err) }
	scotchPos64, err := strconv.ParseInt(strings.TrimSpace(scotchPosLine), 10, 32)
	if err != nil { panic(err) }
	player64, err := strconv.ParseInt(strings.TrimSpace(playerLine), 10, 32)
	if err != nil { panic(err) }
	scotchPos := int(scotchPos64)
	player := int(player64) - 1
	bids := [2][]int{
		toIntSlice(strings.Fields(p1BidsLine)),
		toIntSlice(strings.Fields(p2BidsLine)),
	}
	g := NewGame()
	for i := range(bids[0]) {
		g.next([2]int{bids[0][i], bids[1][i]})
	}
	if g.bottlePos != scotchPos {
		fmt.Println("scotchPos", scotchPos, "g.bottlePos", g.bottlePos)
		panic("bad step")
	}
	playerPos := [2]int{0, 10}
	delta := abs(playerPos[player] - scotchPos)
	if delta == 1 || g.cash[player] == 0 {
		fmt.Println(g.cash[player])
		return
	}
	myBid := g.cash[player] / (delta * 3 / 2)
	if myBid < 1 {
		myBid = 1
	}
	fmt.Println(myBid)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
