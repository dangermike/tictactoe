package mover

import (
	"github.com/dangermike/tictactoe/engine"
)

type Mover interface {
	GetMove(board engine.Board) (x, y uint32)
	Complete(result engine.BoardState)
	Init(mark engine.Mark)
}
