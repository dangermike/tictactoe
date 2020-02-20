package learning

import (
	"math/rand"
	"time"

	"github.com/dangermike/tictactoe/engine"
)

const (
	VALUE_TIE  = int32(1)
	VALUE_WIN  = int32(2)
	VALUE_LOSS = int32(-1)
	MIN_SCORE  = int32(-1 * (1 << 31))
)

type boardPos struct {
	board    engine.Board
	position int
}

type scoreSet [9]int32

type Robot struct {
	rnd          *rand.Rand
	mark         engine.Mark
	NodeScoreSet map[engine.Board]scoreSet
	moves        []boardPos
}

func New() *Robot {
	return &Robot{
		rnd:          rand.New(rand.NewSource(time.Now().UnixNano())),
		NodeScoreSet: map[engine.Board]scoreSet{},
		moves:        make([]boardPos, 9),
	}
}

func (r *Robot) Init(mark engine.Mark) {
	r.mark = mark
	r.moves = r.moves[:0]
}

func (r *Robot) GetMove(board engine.Board) (uint32, uint32) {
	ss, ok := r.NodeScoreSet[board]
	if !ok {
		for i := uint32(0); i < 9; i++ {
			if m, err := board.GetByIndex(i); err == nil && m != engine.MARK_NONE {
				ss[i] = MIN_SCORE
			}
		}
		r.NodeScoreSet[board] = ss
	}

	var moveIx int
	bestScore := MIN_SCORE

	for i := 0; i < len(ss); i++ {
		if ss[i] > MIN_SCORE {
			score := ss[i] + int32(10000*r.rnd.Float64())
			if score > bestScore {
				moveIx = i
				bestScore = score
			}
		}
	}

	r.moves = append(r.moves, boardPos{board, moveIx})
	return uint32(moveIx % 3), uint32(moveIx / 3)
}

func (r *Robot) Complete(result engine.BoardState) {
	var value int32
	switch result {
	case engine.BOARDSTATE_X_WIN:
		if r.mark == engine.MARK_X {
			value = VALUE_WIN
		} else {
			value = VALUE_LOSS
		}
	case engine.BOARDSTATE_O_WIN:
		if r.mark == engine.MARK_O {
			value = VALUE_WIN
		} else {
			value = VALUE_LOSS
		}
	case engine.BOARDSTATE_TIE:
		value = VALUE_TIE
	}
	for _, move := range r.moves {
		ss := r.NodeScoreSet[move.board]
		ss[move.position] += value
		r.NodeScoreSet[move.board] = ss
	}
}
