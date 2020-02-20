package engine

type Game struct {
	board         Board
	currentPlayer Mark
}

func NewGame() *Game {
	return &Game{currentPlayer: MARK_X}
}

func (g *Game) Move(x, y uint32) (Board, error) {
	newBoard, err := g.board.Apply(x, y, g.currentPlayer)
	if err != nil {
		return g.board, err
	}
	g.board = newBoard
	g.currentPlayer ^= MARK_BOTH
	return g.board, nil
}

func (g *Game) State() BoardState {
	return g.board.State()
}

func (g *Game) CurrentPlayer() Mark {
	return g.currentPlayer
}

func (g *Game) Reset() {
	g.board = Board(0)
	g.currentPlayer = MARK_X
}
