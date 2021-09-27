package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/dangermike/tictactoe/engine"
	"github.com/dangermike/tictactoe/player"
	"github.com/dangermike/tictactoe/player/dumb"
	"github.com/dangermike/tictactoe/player/heuristic"
	"github.com/dangermike/tictactoe/player/learning"
	"github.com/dangermike/tictactoe/player/learningminimizing"
)

func main() {
	p1 := learningminimizing.New()
	p2 := learning.New()
	mrDumb := dumb.New()

	matches := map[string][2]player.Player{
		"mvl": [2]player.Player{p1, p2},
		"mvd": [2]player.Player{p1, mrDumb},
		"lvd": [2]player.Player{p2, mrDumb},
		"dvd": [2]player.Player{mrDumb, mrDumb},
	}

	for i := 0; i < 1; i++ {
		for name, players := range matches {
			start := time.Now()
			p1WinRate, p2WinRate, tieRate := fight(players[0], players[1])
			duration := time.Since(start)
			fmt.Printf("%s: %.2f/%.2f/%.2f (%dms)\n", name, p1WinRate, p2WinRate, tieRate, duration.Milliseconds())
		}
		fmt.Println("----------")
	}

	boards := make([]engine.Board, 0, len(p1.NodeScoreSet))
	for board := range p1.NodeScoreSet {
		boards = append(boards, board)
	}
	hplayer := heuristic.New()
	sort.Slice(boards, func(i, j int) bool { return boards[i] < boards[j] })
	for _, board := range boards {
		x, _ := hplayer.GetMove(board)
		if x > 2 {
			fmt.Print("[")
			fmt.Print(board)
			fmt.Println("]")
		}
	}
}

func fight(p1 player.Player, p2 player.Player) (p1WinRate, p2WinRate, tieRate float64) {
	g := engine.NewGame()

	results := [3]int{}
	cnt := 0.0
	for x := 0; x < 100000; x++ {
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

func RunGame(g *engine.Game, player1 player.Player, player2 player.Player) engine.BoardState {
	players := [2]player.Player{player1, player2}
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
