package main

import (
	"fmt"
	"testing"

	"github.com/dangermike/tictactoe/engine"
	"github.com/dangermike/tictactoe/mover/dumb"
)

func BenchmarkDumbGame(b *testing.B) {
	results := map[engine.BoardState]int{}
	p1 := dumb.New()
	p2 := dumb.New()
	g := engine.NewGame()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := RunGame(g, p1, p2)
		results[result] = 1 + results[result]
		g.Reset()
	}
	fmt.Println(results)
}
