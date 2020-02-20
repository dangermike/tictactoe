package learningminimizing

import (
	"fmt"
	"testing"

	"github.com/dangermike/tictactoe/engine"
	"github.com/stretchr/testify/assert"
)

func TestGetMove(t *testing.T) {
	board := makeBoardFromString("X  O  X  ")
	robot := New()
	cnts := [9]int{}
	iterations := 1000000
	for x := 0; x < iterations; x++ {
		x, y := robot.GetMove(board)
		i := x + (3 * y)
		cnts[i]++
	}

	fmt.Println(cnts)

	// tests that we never choose occupied squares
	assert.Equal(t, 0, cnts[0])
	assert.Equal(t, 0, cnts[3])
	assert.Equal(t, 0, cnts[6])

	// tests that we are evenly distributed on the open squares
	exp := int(float64(iterations/6) * 0.8)
	assert.Less(t, exp, cnts[1])
	assert.Less(t, exp, cnts[2])
	assert.Less(t, exp, cnts[4])
	assert.Less(t, exp, cnts[5])
	assert.Less(t, exp, cnts[7])
	assert.Less(t, exp, cnts[8])
}

func makeBoardFromString(b string) engine.Board {
	var board engine.Board
	for ix, s := range b {
		if s == 'x' || s == 'X' {
			board |= engine.Board(engine.MARK_X << (ix << 1))
			continue
		}
		if s == 'o' || s == 'O' {
			board |= engine.Board(engine.MARK_O << (ix << 1))
			continue
		}
	}
	return board
}
