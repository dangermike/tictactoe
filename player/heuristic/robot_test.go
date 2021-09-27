package heuristic_test

import (
	"testing"

	"github.com/dangermike/tictactoe/engine"
	"github.com/dangermike/tictactoe/player/heuristic"
	"github.com/stretchr/testify/assert"
)

func TestForced(t *testing.T) {
	board := engine.FromString("O OX XO X")
	x, y := heuristic.New().GetMove(board)
	assert.Equal(t, uint32(1), x)
	assert.Equal(t, uint32(0), y)
}
