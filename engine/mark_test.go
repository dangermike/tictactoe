package engine_test

import (
	"testing"

	ttt "github.com/dangermike/tictactoe/engine"
	"github.com/stretchr/testify/assert"
)

func TestMarkString(t *testing.T) {
	assert.Equal(t, "X", ttt.MARK_X.String())
	assert.Equal(t, "O", ttt.MARK_O.String())
	assert.Equal(t, "B", ttt.MARK_BOTH.String())
	assert.Equal(t, " ", ttt.MARK_NONE.String())
	assert.Equal(t, "!", ttt.MARK_BAD.String())
	assert.Equal(t, "?", ttt.Mark(123).String())
}
