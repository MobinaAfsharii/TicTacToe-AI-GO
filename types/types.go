package types

import "fmt"

type Marker string

const (
	MarkerX Marker = "X"
	MarkerO Marker = "O"
	Empty   Marker = " "
)

type PlayerType string

const (
	PlayerHuman PlayerType = "Player"
	PlayerAI    PlayerType = "AI"
)

type Board [9]Marker

type Score struct {
	Player int
	AI     int
	Draws  int
}

type WinningCombo [3]int

func ValidMarkers() []Marker {
	return []Marker{MarkerX, MarkerO}
}

func ValidPlayerTypes() []PlayerType {
	return []PlayerType{PlayerHuman, PlayerAI}
}

func (m Marker) IsValid() bool {
	return m == MarkerX || m == MarkerO
}

func (m Marker) OppositeMarker() (Marker, error) {
	switch m {
	case MarkerX:
		return MarkerO, nil
	case MarkerO:
		return MarkerX, nil
	default:
		return Empty, fmt.Errorf("invalid marker: %v", m)
	}
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

var WinningCombinations = []WinningCombo{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
	{0, 4, 8}, {2, 4, 6},
}

func (b Board) CheckWin(marker Marker) (WinningCombo, bool) {
	for _, combo := range WinningCombinations {
		if b[combo[0]] == marker && b[combo[1]] == marker && b[combo[2]] == marker {
			return combo, true
		}
	}
	return WinningCombo{}, false
}
