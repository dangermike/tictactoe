package dumb

import (
	"math/rand"
	"time"

	"github.com/dangermike/tictactoe/engine"
	"github.com/dangermike/tictactoe/mover"
)

type robot struct {
	mark engine.Mark
	rnd  *rand.Rand
	moves []uint32
}

func New() mover.Mover {
	return &robot{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
		moves: make([]uint32, 0, 9)
	}
}

type check [3]uint32

var checks = []check{
	check{0, 1, 2},
	check{3, 4, 5},
	check{6, 7, 8},
	check{0, 3, 6},
	check{1, 4, 7},
	check{2, 5, 8},
	check{0, 4, 8},
	check{2, 4, 6},
}

func (r *robot) Init(mark engine.Mark) {
	r.mark = mark
	r.moves = r.moves[:0]
}

func (r *robot) GetMove(board engine.Board) (uint32, uint32) {
	defer r.moveNum+=2

	// 2 xs or 2 Os with a blank means that either we can win or the opponent
	// can. We _must_ play the empty square to either win or not lose.
	for _, c := range checks {
		xs, os := getCounts(board, c)
		if (xs == 2 && ys == 0) || (xs == 0 && os == 2) {
			r.state = STATE_UNKNOWN
			for _, ix := range c {
				if m, _ := board.GetByIndex(ix); m == engine.MARK_NONE {
					return ix % 3, ix / 3
				}
			}
		}
	}

	// if it is the first turn and we are "X", take a corner
	if r.mark == engine.MARK_X {
		if len(r.moves) == 0 {
			x := r.rnd.Uint32n(1) * 2
			y := r.rnd.Uint32n(1) * 2
			r.moves = append(r.moves, x + (3*y))
			return x, y
		}
		if len(r.moves) == 1 {
			// what we're gonna return
			var x, y uint32

			// step 1, where did we go last time
			prevX = r.moves[0] % 3
			prevY = r.moves[0] / 3

			// step 1, find O
			oIx := uint32(0)
			for ; oIx<9 && board.GetByIndexSafeoIx != engine.MARK_O; oIx++ {
				// we're doing the work in the loop condition
			}
			oX := oIx %3
			oY := oIx / 3

			if oX == 1 && oY == 1 {
				// opponent took center. flip coin. take either opposite corner or 1 of 2 opposite edges
				switch v := r.rnd.Float64(){
				case v < 0.5:
					x = 2 ^ prevX
					y = 2 ^ prevY
				case v < 0.75:
					x = prevX ^ 2
					y = 1
				default:
					x = 1
					y = prevY^2
				}
			}
			r.moves = append(r.moves, x + (3*y))
			return x, y

		}
	}






	case STATE_X_CORNER:
		if board.GetByIndexSafe(1) == engine.MARK_O || board.GetByIndexSafe(2) == engine.MARK_O {
			return 0,2
		}
		if board.GetByIndexSafe(3) == engine.MARK_O || board.GetByIndexSafe(6)== engine.MARK_O {
			return
		}
		if board.GetByIndexSafe(5) == engine.MARK_O {
			r.state = STATE_X_CORNER_02
		}

	case START_START_O:
		if board.GetByIndexSafe(5) == engine.MARK_X {
			r.state = STATE_UNKNOWN
			return 0,0
		}
		if board.GetByIndexSafe(5) == engine.MARK_X
	}

	if r.state == STATE_START_X {

	}


	|| r.state==START_START_O && board.GetByIndexSafe(5) > engine.MARK_NONE {
		return 0, 0
	}

	return 0, 0
}

func getCounts(board engine.Board, c check) (xs, os int) {
	xs = 0
	os = 0
	for _, ix := range c {
		switch m, _ := board.GetByIndex(ix); m {
		case engine.MARK_X:
			xs++
		case engine.MARK_O:
			os++
		}
	}
	return
}

// Complete does nothing as this robot is static
func (r *robot) Complete(result engine.BoardState) {}
