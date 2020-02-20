package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dangermike/tictactoe/engine"
	"github.com/dangermike/tictactoe/mover"
)

type human struct {
	mark engine.Mark
}

func New() mover.Mover {
	return &human{}
}

func (h *human) Init(mark engine.Mark) {
	h.mark = mark
	fmt.Printf("You are playing as %s\n", h.mark.String())
}

func (h *human) GetMove(board engine.Board) (x, y uint32) {
	for {
		printBoard(board)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Your move, %s: ", h.mark.String())
		if !scanner.Scan() {
			panic(scanner.Err())
		}
		x, y, ok := parseMove(scanner.Text())
		if !ok {
			fmt.Println("Please provide a comma-delimted, 0-based coordinate")
			continue
		}
		return x, y
	}
}

func (h *human) Complete(result engine.BoardState) {
	fmt.Printf("The game is over. %s\n", result.String())
}

func parseMove(line string) (x, y uint32, ok bool) {
	s := strings.Split(line, ",")
	if len(s) != 2 {
		return 0, 0, false
	}
	xv, err := strconv.ParseUint(s[0], 10, 32)
	if xv > 2 || err != nil {
		return 0, 0, false
	}
	yv, err := strconv.ParseUint(s[1], 10, 32)
	if yv > 2 || err != nil {
		return 0, 0, false
	}
	return uint32(xv), uint32(yv), true
}

func printBoard(board engine.Board) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			fmt.Print(" ")
			fmt.Print(board.GetSafe(uint32(x), uint32(y)))
			fmt.Print(" ")
			if x < 2 {
				fmt.Print("│")
			}
		}
		fmt.Println()
		if y < 2 {
			fmt.Println("───┼───┼───")
		}
	}
}
