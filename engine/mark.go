package engine

type Mark uint32

const (
	MARK_NONE = Mark(0)
	MARK_X    = Mark(1)
	MARK_O    = Mark(2)
	MARK_BOTH = MARK_X | MARK_O
	MARK_BAD  = Mark(1 << 31)
)

func (m Mark) String() string {
	switch m {
	case MARK_NONE:
		return " "
	case MARK_X:
		return "X"
	case MARK_O:
		return "O"
	case MARK_BOTH:
		return "B"
	case MARK_BAD:
		return "!"
	default:
		return "?"
	}
}
