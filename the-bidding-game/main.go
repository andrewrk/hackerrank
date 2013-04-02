package main

import "fmt"

type Game struct {
	playerCash [2]int
	playerState [2]int
	drawAdvantage int
}

var otherPlayer = [2]int{1, 0}
var playerName = [2]string{"P1", "P2"}
const PRINT = false

func NewGame() Game {
	return Game{[2]int{25, 25}, [2]int{2, 2}, 0}
}

func (g * Game) next(bets [2]int) {
	if PRINT {
		fmt.Println("P1State", g.playerState[0], "P2State", g.playerState[1], "drawAdv", playerName[g.drawAdvantage], "P1Cash", g.playerCash[0], "P2Cash", g.playerCash[1], "P1Bid", bets[0], "P2Bid", bets[1])
	}
	var winner int
	if bets[0] > bets[1] {
		winner = 0
	} else if bets[1] > bets[0] {
		winner = 1
	} else {
		winner = g.drawAdvantage
		g.drawAdvantage = otherPlayer[winner]
	}
	g.playerCash[winner] -= bets[winner]
	g.playerState[winner] -= 1
}

func (g * Game) computeWinner() int {
	// the player who has not moved will try to minimally
	// match the other player's bet
	loser := 0
	if g.playerState[0] < g.playerState[1] { loser = 1 }
	winner := otherPlayer[loser]
	winnerBet := g.playerCash[winner]
	loserBet := winnerBet
	if g.drawAdvantage != loser { loserBet += 1 }

	// we already know that the winner of the first bet
	// won the game
	if loserBet > g.playerCash[loser] {
		return winner
	}

	// need to do another round of bidding to find out
	var bets [2]int
	bets[loser] = loserBet
	bets[winner] = winnerBet
	g.next(bets)

	if g.playerState[0] == 0 { return 0 }
	if g.playerState[1] == 0 { return 1 }

	if g.playerState[0] != 1 && g.playerState[1] != 1 {
		if g.playerState[0] == 2 && g.playerState[1] == 2 {
			if g.playerCash[0] > g.playerCash[1] {
				return 0
			} else if g.playerCash[1] > g.playerCash[0] {
				return 1
			} else {
				return g.drawAdvantage
			}
		} else {
			fmt.Println(g)
			panic("bad step")
		}
	}

	// both players bet everything they have
	g.next(g.playerCash)

	if g.playerState[0] == 0 { return 0 }
	if g.playerState[1] == 0 { return 1 }

	// both players bet everything they have
	g.next(g.playerCash)

	if g.playerState[0] == 0 { return 0 }
	if g.playerState[1] == 0 { return 1 }

	panic("bad step")
}

func (originalGame Game) computeWinner2() int {
	var wins [2][101]int
	for p1bet := 1; p1bet <= originalGame.playerCash[0]; p1bet++ {
		for p2bet := 1; p2bet <= originalGame.playerCash[1]; p2bet++ {
			// set up this game
			g := originalGame
			// do the first bet
			g.next([2]int{p1bet, p2bet})
			// from there we can figure out who wins
			winner := g.computeWinner()
			winnerBet := p1bet
			if winner == 1 { winnerBet = p2bet }
			fmt.Print(winner)
			wins[winner][winnerBet] += 1
			//fmt.Println("P1Bet", p1bet, "P2Bet", p2bet, "winner", playerName[winner])
		}
		fmt.Print("\n")
	}
	for i := 0; i < 2; i++ {
		for bet := 1; bet <= originalGame.playerCash[otherPlayer[i]]; bet++ {
			if wins[i][bet] == originalGame.playerCash[otherPlayer[i]] {
				return i
			}
		}
	}
	return 2
}

func main() {
	// set up every possible game
	originalGame := NewGame()
	originalGame.computeWinner2()
	return
	for p1bet := 1; p1bet <= originalGame.playerCash[0]; p1bet++ {
		for p2bet := 1; p2bet <= originalGame.playerCash[1]; p2bet++ {
			g := originalGame
			g.next([2]int{p1bet, p2bet})
			w := g.computeWinner2()
			fmt.Print(w)
		}
		fmt.Print("\n")
	}
}
