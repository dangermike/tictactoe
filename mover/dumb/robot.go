package dumb

import (
	"math/rand"
	"time"

	"github.com/dangermike/tictactoe/engine"
	"github.com/dangermike/tictactoe/mover"
)

type pos struct {
	x uint32
	y uint32
}

type robot struct {
	rnd   *rand.Rand
	moves []pos
}

func New() mover.Mover {
	return &robot{
		rnd:   rand.New(rand.NewSource(time.Now().UnixNano())),
		moves: make([]pos, 0, 9),
	}
}

func (r *robot) Init(mark engine.Mark) {}

func (r *robot) GetMove(board engine.Board) (x, y uint32) {
	r.moves = r.moves[:0]
	for x := uint32(0); x < 3; x++ {
		for y := uint32(0); y < 3; y++ {
			if m, err := board.Get(x, y); err == nil && m == engine.MARK_NONE {
				r.moves = append(r.moves, pos{x, y})
			}
		}
	}
	moveIx := r.rnd.Intn(len(r.moves))
	return r.moves[moveIx].x, r.moves[moveIx].y
}

func (r *robot) Complete(result engine.BoardState) {}
