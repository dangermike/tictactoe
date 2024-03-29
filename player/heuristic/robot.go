package heuristic

import (
	"math/rand"
	"time"

	"github.com/dangermike/tictactoe/engine"
	"github.com/dangermike/tictactoe/player"
)

const NOMOVE = uint32(1 << 31)

var moveMap = map[engine.Board][]uint32{
	engine.FromString("         "): []uint32{},
	engine.FromString("X        "): []uint32{},
	engine.FromString(" X       "): []uint32{},
	engine.FromString("OX       "): []uint32{},
	engine.FromString("XO       "): []uint32{},
	engine.FromString("O X      "): []uint32{},
	engine.FromString("OXX      "): []uint32{},
	engine.FromString("XOX      "): []uint32{},
	engine.FromString("OX X     "): []uint32{},
	engine.FromString(" O X     "): []uint32{},
	engine.FromString("O XX     "): []uint32{},
	engine.FromString(" OXX     "): []uint32{},
	engine.FromString("OOXX     "): []uint32{},
	engine.FromString("  OX     "): []uint32{},
	engine.FromString(" XOX     "): []uint32{},
	engine.FromString("OXOX     "): []uint32{},
	engine.FromString(" X O     "): []uint32{},
	engine.FromString("  XO     "): []uint32{},
	engine.FromString("XOXO     "): []uint32{},
	engine.FromString("XXOO     "): []uint32{},
	engine.FromString("    X    "): []uint32{},
	engine.FromString("O   X    "): []uint32{},
	engine.FromString(" O  X    "): []uint32{},
	engine.FromString("X   O    "): []uint32{},
	engine.FromString(" X  O    "): []uint32{},
	engine.FromString(" X XO    "): []uint32{},
	engine.FromString("  XXO    "): []uint32{},
	engine.FromString("   O X   "): []uint32{},
	engine.FromString("X  O X   "): []uint32{},
	engine.FromString(" X O X   "): []uint32{},
	engine.FromString("XO O X   "): []uint32{},
	engine.FromString("X OO X   "): []uint32{},
	engine.FromString(" XOO X   "): []uint32{},
	engine.FromString("XXOO X   "): []uint32{},
	engine.FromString("   OXX   "): []uint32{},
	engine.FromString(" O OXX   "): []uint32{},
	engine.FromString("  OOXX   "): []uint32{},
	engine.FromString("   XOX   "): []uint32{},
	engine.FromString("X  OOX   "): []uint32{},
	engine.FromString(" X OOX   "): []uint32{},
	engine.FromString("  O   X  "): []uint32{},
	engine.FromString(" XO   X  "): []uint32{},
	engine.FromString("OXO   X  "): []uint32{},
	engine.FromString("OXOX  X  "): []uint32{},
	engine.FromString("X OO  X  "): []uint32{},
	engine.FromString(" XOO  X  "): []uint32{},
	engine.FromString("XXOO  X  "): []uint32{},
	engine.FromString("  O X X  "): []uint32{},
	engine.FromString("  OOX X  "): []uint32{},
	engine.FromString("  X O X  "): []uint32{},
	engine.FromString(" XO O X  "): []uint32{},
	engine.FromString("O    XX  "): []uint32{},
	engine.FromString(" O   XX  "): []uint32{},
	engine.FromString("  O  XX  "): []uint32{},
	engine.FromString("OXO  XX  "): []uint32{},
	engine.FromString("O  O XX  "): []uint32{},
	engine.FromString("OX O XX  "): []uint32{},
	engine.FromString(" O O XX  "): []uint32{},
	engine.FromString("XO O XX  "): []uint32{},
	engine.FromString("  OO XX  "): []uint32{},
	engine.FromString("X OO XX  "): []uint32{},
	engine.FromString(" XOO XX  "): []uint32{},
	engine.FromString("OXOO XX  "): []uint32{},
	engine.FromString("XOOO XX  "): []uint32{},
	engine.FromString("  OOXXX  "): []uint32{},
	engine.FromString("  O OXX  "): []uint32{},
	engine.FromString(" XO OXX  "): []uint32{},
	engine.FromString("X  OOXX  "): []uint32{},
	engine.FromString(" X OOXX  "): []uint32{},
	engine.FromString("X OOOXX  "): []uint32{},
	engine.FromString(" XOOOXX  "): []uint32{},
	engine.FromString("XXOOOXX  "): []uint32{},
	engine.FromString(" X   OX  "): []uint32{},
	engine.FromString("OX   OX  "): []uint32{},
	engine.FromString("O  X OX  "): []uint32{},
	engine.FromString("OX X OX  "): []uint32{},
	engine.FromString("XOX   O  "): []uint32{},
	engine.FromString("O XX  O  "): []uint32{},
	engine.FromString("OXXX  O  "): []uint32{},
	engine.FromString(" OXX  O  "): []uint32{},
	engine.FromString("XOXX  O  "): []uint32{},
	engine.FromString(" OX X O  "): []uint32{},
	engine.FromString(" X   XO  "): []uint32{},
	engine.FromString("XO   XO  "): []uint32{},
	engine.FromString(" X X OO  "): []uint32{},
	engine.FromString("OXXX OO  "): []uint32{},
	engine.FromString("XOXX OO  "): []uint32{},
	engine.FromString("O XXXOO  "): []uint32{},
	engine.FromString(" OXXXOO  "): []uint32{},
	engine.FromString("OOXX   X "): []uint32{},
	engine.FromString(" OXO   X "): []uint32{},
	engine.FromString("XOXO   X "): []uint32{},
	engine.FromString("X OO   X "): []uint32{},
	engine.FromString(" OXXO  X "): []uint32{},
	engine.FromString(" O O X X "): []uint32{},
	engine.FromString("XO O X X "): []uint32{},
	engine.FromString("X OO X X "): []uint32{},
	engine.FromString("XOOO X X "): []uint32{},
	engine.FromString(" O OXX X "): []uint32{},
	engine.FromString(" O XOX X "): []uint32{},
	engine.FromString("XO OOX X "): []uint32{},
	engine.FromString(" OX   OX "): []uint32{},
	engine.FromString("XOX   OX "): []uint32{},
	engine.FromString("O XX  OX "): []uint32{},
	engine.FromString(" OXX  OX "): []uint32{},
	engine.FromString("OOXX  OX "): []uint32{},
	engine.FromString("XOXO  OX "): []uint32{},
	engine.FromString(" OX X OX "): []uint32{},
	engine.FromString("XOX O OX "): []uint32{},
	engine.FromString(" OXXO OX "): []uint32{},
	engine.FromString("XOXXO OX "): []uint32{},
	engine.FromString("XOX  OOX "): []uint32{},
	engine.FromString("XO X OOX "): []uint32{},
	engine.FromString("O XX OOX "): []uint32{},
	engine.FromString(" OXX OOX "): []uint32{},
	engine.FromString("XOXX OOX "): []uint32{},
	engine.FromString(" OX XOOX "): []uint32{},
	engine.FromString(" O XXOOX "): []uint32{},
	engine.FromString(" OXXXOOX "): []uint32{},
	engine.FromString("XXOO X O "): []uint32{},
	engine.FromString("OXOX  XO "): []uint32{},
	engine.FromString("XXOO  XO "): []uint32{},
	engine.FromString("OXO X XO "): []uint32{},
	engine.FromString("OXO  XXO "): []uint32{},
	engine.FromString("X OO XXO "): []uint32{},
	engine.FromString("XXOO XXO "): []uint32{},
	engine.FromString("O XX OO X"): []uint32{},
	engine.FromString("OXXX OO X"): []uint32{},
	engine.FromString("O XXXOO X"): []uint32{},
}

type robot struct {
	mark  engine.Mark
	rnd   *rand.Rand
	moves []uint32
}

func New() player.Player {
	return &robot{
		rnd:   rand.New(rand.NewSource(time.Now().UnixNano())),
		moves: make([]uint32, 0, 9),
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
	// 2 xs or 2 Os with a blank means that either we can win or the opponent
	// can. We _must_ play the empty square to either win or not lose.
	for _, c := range checks {
		xs, os := getCounts(board, c)
		if (xs == 2 && os == 0) || (xs == 0 && os == 2) {
			for _, ix := range c {
				if m, _ := board.GetByIndex(ix); m == engine.MARK_NONE {
					return ix % 3, ix / 3
				}
			}
		}
	}

	// // if it is the first turn and we are "X", take a corner
	// if r.mark == engine.MARK_X {
	// 	if len(r.moves) == 0 {
	// 		x := (r.rnd.Uint32() % 2) * 2
	// 		y := (r.rnd.Uint32() % 2) * 2
	// 		r.moves = append(r.moves, x+(3*y))
	// 		return x, y
	// 	}
	// 	if len(r.moves) == 1 {
	// 		// what we're gonna return
	// 		var x, y uint32

	// 		// step 1, where did we go last time
	// 		prevX := r.moves[0] % 3
	// 		prevY := r.moves[0] / 3

	// 		// step 1, find O
	// 		oIx := uint32(0)
	// 		for ; oIx < 9 && board.GetByIndexSafeoIx != engine.MARK_O; oIx++ {
	// 			// we're doing the work in the loop condition
	// 		}
	// 		oX := oIx % 3
	// 		oY := oIx / 3

	// 		if oX == 1 && oY == 1 {
	// 			// opponent took center. flip coin. take either opposite corner or 1 of 2 opposite edges
	// 			switch v := r.rnd.Float64(); v {
	// 			case v < 0.5:
	// 				x = 2 ^ prevX
	// 				y = 2 ^ prevY
	// 			case v < 0.75:
	// 				x = prevX ^ 2
	// 				y = 1
	// 			default:
	// 				x = 1
	// 				y = prevY ^ 2
	// 			}
	// 		}
	// 		r.moves = append(r.moves, x+(3*y))
	// 		return x, y

	// 	}
	// }

	return NOMOVE, NOMOVE
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
