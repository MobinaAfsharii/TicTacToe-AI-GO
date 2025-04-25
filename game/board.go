package game

import (
	"tic-tac-toe/types"
)

func NewBoard() types.Board {
	return types.Board{types.Empty, types.Empty, types.Empty, types.Empty, types.Empty, types.Empty, types.Empty, types.Empty, types.Empty}
}
