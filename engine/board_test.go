package engine_test

import (
	"testing"

	ttt "github.com/dangermike/tictactoe/engine"
	"github.com/stretchr/testify/assert"
)

func makeBoardFromString(b string) ttt.Board {
	var board ttt.Board
	for ix, s := range b {
		if s == 'x' || s == 'X' {
			board |= ttt.Board(ttt.MARK_X << (ix << 1))
			continue
		}
		if s == 'o' || s == 'O' {
			board |= ttt.Board(ttt.MARK_O << (ix << 1))
			continue
		}
	}
	return board
}

func TestBoardApply(t *testing.T) {
	var board ttt.Board
	board, err := board.Apply(1, 0, ttt.MARK_O)
	assert.NoError(t, err)
	board = board.ApplySafe(1, 1, ttt.MARK_O)
	mark, err := board.Get(1, 0)
	assert.NoError(t, err)
	assert.Equal(t, ttt.MARK_O, mark)

	_, err = board.Apply(1, 0, ttt.MARK_O)
	assert.Error(t, err)

	_, err = board.Apply(3, 0, ttt.MARK_O)
	assert.Error(t, err)

	_, err = board.Apply(0, 0, ttt.Mark(123))
	assert.Error(t, err)

	_, err = board.Get(3, 0)
	assert.Error(t, err)
}

func TestBoardStateOpen(t *testing.T) {
	var board ttt.Board
	assert.Equal(t, ttt.BOARDSTATE_OPEN, board.State())
}

func TestBoardState(t *testing.T) {
	// X O X
	// X O X
	// O X O
	assert.Equal(t, ttt.BOARDSTATE_TIE, makeBoardFromString("XOXXOXOXO").State())

	// X O X
	// X O X
	// O X
	assert.Equal(t, ttt.BOARDSTATE_OPEN, makeBoardFromString("XOXXOXOX ").State())

	// X O X
	// X O X
	// O X X
	assert.Equal(t, ttt.BOARDSTATE_X_WIN, makeBoardFromString("XOXXOXOXX").State())

	// X O X
	// X O X
	// O O
	assert.Equal(t, ttt.BOARDSTATE_O_WIN, makeBoardFromString("XOXXOXOO ").State())

	// O O X
	// X O X
	// O X O
	assert.Equal(t, ttt.BOARDSTATE_O_WIN, makeBoardFromString("OOXXOXOXO").State())

	// X O X
	// O X X
	// X X O
	assert.Equal(t, ttt.BOARDSTATE_X_WIN, makeBoardFromString("XOXOXXXXO").State())

	// X O
	// O X X
	// O X O
	assert.Equal(t, ttt.BOARDSTATE_TIE, makeBoardFromString("XO OXXOXO").State())
}

func TestBoardStateString(t *testing.T) {
	assert.Equal(t, "open", ttt.BOARDSTATE_OPEN.String())
	assert.Equal(t, "O wins", ttt.BOARDSTATE_O_WIN.String())
	assert.Equal(t, "X wins", ttt.BOARDSTATE_X_WIN.String())
	assert.Equal(t, "tie", ttt.BOARDSTATE_TIE.String())
	assert.Equal(t, "UNKNOWN", ttt.BoardState(123).String())
}

func TestBoardString(t *testing.T) {
	assert.Equal(t, "XOX/OXO/XOX", makeBoardFromString("XOXOXOXOX").String())
	assert.Equal(t, "   /   /   ", makeBoardFromString("").String())
}

func BenchmarkBoardState(b *testing.B) {
	board := makeBoardFromString("XOXXOXOXO")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if ttt.BOARDSTATE_TIE != board.State() {
			panic("boardstate mismatch")
		}
	}
}

func RecTestState(t *testing.T, board ttt.Board, loc uint32) int {
	if loc > 8 {
		return 1
	}
	x := loc % 3
	y := loc / 3

	cnt := 0
	cnt += RecTestState(t, board, loc+1)
	cnt += RecTestState(t, board.ApplySafe(x, y, ttt.MARK_X), loc+1)
	cnt += RecTestState(t, board.ApplySafe(x, y, ttt.MARK_O), loc+1)
	return cnt
}

func TestAll(t *testing.T) {
	var board ttt.Board
	assert.Equal(t, 19683, RecTestState(t, board, 0))
}

func TestGetByIndex(t *testing.T) {
	mark, err := makeBoardFromString("XO O X XO").GetByIndex(3)
	assert.Equal(t, ttt.MARK_O, mark)
	assert.NoError(t, err)

	mark, err = makeBoardFromString("XO O X XO").GetByIndex(22)
	assert.Equal(t, ttt.MARK_BAD, mark)
	assert.Equal(t, ttt.ErrInvalidCoordinate, err)
}

func TestRotate(t *testing.T) {
	board := makeBoardFromString("XO O X XO")
	assert.Equal(t, board, board.Rotate(ttt.ROT_IDENTITY))
	assert.NotEqual(t, board, board.Rotate(ttt.ROT_RIGHT))

	assert.Equal(
		t,
		board,
		board.Rotate(ttt.ROT_RIGHT).Rotate(ttt.ROT_LEFT),
	)

	assert.Equal(
		t,
		board,
		board.Rotate(ttt.ROT_180).Rotate(ttt.ROT_180),
	)

	assert.Equal(
		t,
		board,
		board.Rotate(ttt.FLIP_H).Rotate(ttt.FLIP_H),
	)

	assert.Equal(
		t,
		board,
		board.Rotate(ttt.FLIP_V).Rotate(ttt.FLIP_V),
	)

	assert.Equal(
		t,
		board,
		board.Rotate(ttt.ROT_RIGHT).Rotate(ttt.ROT_RIGHT).Rotate(ttt.ROT_RIGHT).Rotate(ttt.ROT_RIGHT),
	)
}
func TestRotateAddition(t *testing.T) {
	assert.Equal(
		t,
		ttt.ROT_IDENTITY,
		ttt.ROT_IDENTITY.Add(ttt.ROT_IDENTITY),
	)
	assert.Equal(
		t,
		ttt.ROT_IDENTITY,
		ttt.ROT_RIGHT.Add(ttt.ROT_LEFT),
	)
	assert.Equal(
		t,
		ttt.ROT_IDENTITY,
		ttt.ROT_RIGHT.Add(ttt.ROT_RIGHT).Add(ttt.ROT_RIGHT).Add(ttt.ROT_RIGHT),
	)
	assert.Equal(
		t,
		ttt.ROT_IDENTITY,
		ttt.FLIP_H.Add(ttt.FLIP_H),
	)
	assert.Equal(
		t,
		ttt.ROT_IDENTITY,
		ttt.FLIP_H.Add(ttt.ROT_RIGHT).Add(ttt.FLIP_V).Add(ttt.ROT_LEFT),
	)
	assert.NotEqual(
		t,
		ttt.ROT_RIGHT,
		ttt.ROT_RIGHT.Invert(),
	)
	assert.Equal(
		t,
		ttt.ROT_LEFT,
		ttt.ROT_RIGHT.Invert(),
	)
	assert.Equal(
		t,
		ttt.ROT_RIGHT,
		ttt.ROT_RIGHT.Invert().Invert(),
	)
	assert.Equal(
		t,
		ttt.ROT_IDENTITY,
		ttt.ROT_IDENTITY.Invert(),
	)
	assert.Equal(
		t,
		ttt.ROT_IDENTITY,
		ttt.ROT_180.Add(ttt.ROT_180),
	)
}

func TestMinimizeCorner(t *testing.T) {
	board := makeBoardFromString("        X")
	expected := makeBoardFromString("X        ")
	minimized, transform := board.Minimize()
	assert.Equal(t, expected, minimized)
	assert.Equal(t, ttt.ROT_180, transform)
}

func TestMinimizeCorners(t *testing.T) {
	board := makeBoardFromString("O       X")
	expected := makeBoardFromString("  O   X  ")
	minimized, transform := board.Minimize()
	assert.Equal(t, expected.String(), minimized.String())
	assert.Equal(t, ttt.ROT_RIGHT, transform)
}
