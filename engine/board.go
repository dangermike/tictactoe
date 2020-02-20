package engine

import (
	"errors"
)

var (
	ErrInvalidCoordinate = errors.New("invalid coordinate")
	ErrInvalidMarker     = errors.New("invalid marker")
	ErrOccupied          = errors.New("board cell is occupied")
)

type BoardState byte

func (bs BoardState) String() string {
	switch bs {
	case BOARDSTATE_OPEN:
		return "open"
	case BOARDSTATE_O_WIN:
		return "O wins"
	case BOARDSTATE_X_WIN:
		return "X wins"
	case BOARDSTATE_TIE:
		return "tie"
	default:
		return "UNKNOWN"
	}
}

const (
	BOARDSTATE_OPEN  = BoardState(0)
	BOARDSTATE_X_WIN = BoardState(1)
	BOARDSTATE_O_WIN = BoardState(2)
	BOARDSTATE_TIE   = BoardState(3)
)

type Board uint32

func getIx(x, y uint32) (uint32, error) {
	if x > 2 || y > 2 {
		return 0, ErrInvalidCoordinate
	}
	return ((3 * y) + x) << 1, nil
}

func (b Board) GetByIndex(ix uint32) (Mark, error) {
	if ix > 9 {
		return MARK_BAD, ErrInvalidCoordinate
	}
	ix <<= 1 // 2 bits per value
	return Mark(uint32(b)>>ix) & MARK_BOTH, nil

}

func (b Board) GetByIndexSafe(ix uint32) Mark {
	m, err := b.GetByIndex(ix)
	if err != nil {
		panic(err)
	}
	return m
}

func (b Board) Get(x, y uint32) (Mark, error) {
	ix, err := getIx(x, y)
	if err != nil {
		return MARK_BAD, err
	}
	return Mark(uint32(b)>>ix) & MARK_BOTH, nil
}

func (b Board) GetSafe(x, y uint32) Mark {
	m, err := b.Get(x, y)
	if err != nil {
		panic(err)
	}
	return m
}

// Apply puts a value on a board.
// Note that Board is immutable and this makes a copy
func (b Board) Apply(x, y uint32, mark Mark) (Board, error) {
	ix, err := getIx(x, y)
	if err != nil {
		return b, err
	}
	if mark >= MARK_BOTH {
		return b, ErrInvalidMarker
	}

	if Mark(uint32(b)>>ix)&MARK_BOTH > 0 {
		return b, ErrOccupied
	}

	return Board(mark<<ix) | b, nil
}

func (b Board) ApplySafe(x, y uint32, mark Mark) Board {
	b, err := b.Apply(x, y, mark)
	if err != nil {
		panic(err)
	}
	return b
}

type check [3]uint32

func (c check) Check(b Board) BoardState {
	xs := 0
	os := 0
	for _, p := range c {
		m, _ := b.GetByIndex(p)
		if m == MARK_X {
			xs++
		} else if m == MARK_O {
			os++
		}
	}
	if xs == len(c) {
		return BOARDSTATE_X_WIN
	}
	if os == len(c) {
		return BOARDSTATE_O_WIN
	}
	if xs > 0 && os > 0 {
		return BOARDSTATE_TIE
	}
	return BOARDSTATE_OPEN
}

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

func (b Board) State() BoardState {
	isOpen := false
	for _, c := range checks {
		cstate := c.Check(b)
		if cstate == BOARDSTATE_X_WIN || cstate == BOARDSTATE_O_WIN {
			return cstate
		}
		if cstate == BOARDSTATE_OPEN {
			isOpen = true
		}
	}
	if isOpen {
		return BOARDSTATE_OPEN
	}
	return BOARDSTATE_TIE
}

func (b Board) String() string {
	chars := make([]byte, 11)
	chars[3] = '/'
	chars[7] = '/'
	chars[0] = b.GetSafe(0, 0).String()[0]
	chars[1] = b.GetSafe(1, 0).String()[0]
	chars[2] = b.GetSafe(2, 0).String()[0]
	chars[4] = b.GetSafe(0, 1).String()[0]
	chars[5] = b.GetSafe(1, 1).String()[0]
	chars[6] = b.GetSafe(2, 1).String()[0]
	chars[8] = b.GetSafe(0, 2).String()[0]
	chars[9] = b.GetSafe(1, 2).String()[0]
	chars[10] = b.GetSafe(2, 2).String()[0]
	return string(chars)
}

type Rotation [9]uint8

var (
	ROT_IDENTITY = Rotation{0, 1, 2, 3, 4, 5, 6, 7, 8}
	ROT_RIGHT    = Rotation{6, 3, 0, 7, 4, 1, 8, 5, 2}
	ROT_LEFT     = Rotation{2, 5, 8, 1, 4, 7, 0, 3, 6}
	ROT_180      = Rotation{8, 7, 6, 5, 4, 3, 2, 1, 0}
	FLIP_H       = Rotation{2, 1, 0, 5, 4, 3, 8, 7, 6}
	FLIP_V       = Rotation{6, 7, 8, 3, 4, 5, 0, 1, 2}
	allRotations = []Rotation{
		ROT_IDENTITY,
		ROT_RIGHT,
		ROT_LEFT,
		ROT_180,
		FLIP_H,
		FLIP_V,
	}
)

func (r Rotation) Add(other Rotation) Rotation {
	return Rotation{
		r[other[0]],
		r[other[1]],
		r[other[2]],
		r[other[3]],
		r[other[4]],
		r[other[5]],
		r[other[6]],
		r[other[7]],
		r[other[8]],
	}
}

func (r Rotation) Invert() Rotation {
	newR := Rotation{}
	for a, b := range r {
		newR[b] = uint8(a)
	}
	return newR
}

func (b Board) Rotate(rotation Rotation) Board {
	var newBoard Board
	for dIx, sIx := range rotation {
		newBoard |= ((b >> (sIx << 1)) & 3) << (dIx << 1)
	}
	return newBoard
}

func (b Board) Minimize() (Board, Rotation) {
	lastBoard := Board(1<<18 - 1)
	transform := ROT_IDENTITY

	for b < lastBoard {
		lastBoard = b
		newTransform := transform
		newB := b
		for _, r := range allRotations {
			candidate := b.Rotate(r)
			if candidate < newB {
				newB = candidate
				newTransform = transform.Add(r)
			}
		}
		b = newB
		transform = newTransform
	}
	return b, transform
}
