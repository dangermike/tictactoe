package main

import (
	"fmt"

	"github.com/dangermike/tictactoe/engine"
	"github.com/dangermike/tictactoe/mover"
	"github.com/dangermike/tictactoe/mover/dumb"
	"github.com/dangermike/tictactoe/mover/learning"
	"github.com/dangermike/tictactoe/mover/learningminimizing"
)

func main() {
	p1 := learningminimizing.New()
	p2 := learning.New()
	mrDumb := dumb.New()
	for i := 0; i < 50; i++ {
		// RunGame(engine.NewGame(), p1, terminal.New())
		// RunGame(engine.NewGame(), terminal.New(), p1)
		// fmt.Printf("lv[]: xx/xx/xx | %d known boards\n", len(p1.NodeScoreSet))
		p1WinRate, p2WinRate, tieRate := fight(p1, p2)
		fmt.Printf("mvl: %.2f/%.2f/%.2f | %d/%d known boards\n", p1WinRate, p2WinRate, tieRate, len(p1.NodeScoreSet), len(p2.NodeScoreSet))
		p1WinRate, p2WinRate, tieRate = fight(p1, mrDumb)
		fmt.Printf("mvd: %.2f/%.2f/%.2f | %d known boards\n", p1WinRate, p2WinRate, tieRate, len(p1.NodeScoreSet))
		p1WinRate, p2WinRate, tieRate = fight(p2, mrDumb)
		fmt.Printf("lvd: %.2f/%.2f/%.2f | %d known boards\n", p1WinRate, p2WinRate, tieRate, len(p2.NodeScoreSet))
		fmt.Println("----------")
	}
	// for board, scores := range p1.NodeScoreSet {
	// 	high := 0
	// 	for i := 1; i < 9; i++ {
	// 		if scores[i] > scores[high] {
	// 			high = i
	// 		}
	// 	}
	// 	fmt.Printf("%s: %d\n", board, high)
	// }
}

func fight(p1 mover.Mover, p2 mover.Mover) (p1WinRate, p2WinRate, tieRate float64) {
	g := engine.NewGame()

	results := [3]int{}
	cnt := 0.0
	for x := 0; x < 1000000; x++ {
		cnt++
		g.Reset()
		var result engine.BoardState
		if x%2 == 0 {
			result = RunGame(g, p1, p2)
			switch result {
			case engine.BOARDSTATE_X_WIN:
				results[0] += 1
			case engine.BOARDSTATE_O_WIN:
				results[1] += 1
			default:
				results[2] += 1
			}
		} else {
			result = RunGame(g, p2, p1)
			switch result {
			case engine.BOARDSTATE_O_WIN:
				results[0] += 1
			case engine.BOARDSTATE_X_WIN:
				results[1] += 1
			default:
				results[2] += 1
			}
		}
	}
	return float64(results[0]) / cnt,
		float64(results[1]) / cnt,
		float64(results[2]) / cnt
	// fmt.Println(len(p1.NodeScoreSet))
}

func RunGame(g *engine.Game, player1 mover.Mover, player2 mover.Mover) engine.BoardState {
	players := [2]mover.Mover{player1, player2}
	var board engine.Board
	playerIx := 0
	player1.Init(engine.MARK_X)
	player2.Init(engine.MARK_O)
	for g.State() == engine.BOARDSTATE_OPEN {
		player := players[playerIx]
		x, y := player.GetMove(board)
		var err error
		if board, err = g.Move(x, y); err != nil {
			fmt.Println(err)
			continue
		}
		playerIx = (playerIx + 1) % 2
	}
	state := g.State()
	player1.Complete(state)
	player2.Complete(state)
	return state
}
