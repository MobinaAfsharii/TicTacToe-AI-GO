package game

import (
	"fmt"
)


var WinningCombinations = []WinningCombo{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, 
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, 
	{0, 4, 8}, {2, 4, 6}, 
}


func NewBoard() Board {
	return Board{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}
}


func (b Board) IsFull() bool {
	for _, cell := range b {
		if cell == Empty {
			return false
		}
	}
	return true
}


func (b Board) AvailableMoves() []int {
	moves := make([]int, 0, 9)
	for i, cell := range b {
		if cell == Empty {
			moves = append(moves, i)
		}
	}
	return moves
}


func (b *Board) MakeMove(index int, marker Marker) error {
	if index < 0 || index >= 9 {
		return fmt.Errorf("index out of range: %d", index)
	}
	if b[index] != Empty {
		return fmt.Errorf("cell %d is not empty", index)
	}
	if !marker.IsValid() {
		return fmt.Errorf("invalid marker: %v", marker)
	}
	b[index] = marker
	return nil
}


func (b Board) CheckWin(marker Marker) (WinningCombo, bool) {
	for _, combo := range WinningCombinations {
		if b[combo[0]] == marker && b[combo[1]] == marker && b[combo[2]] == marker {
			return combo, true
		}
	}
	return WinningCombo{}, false
}
