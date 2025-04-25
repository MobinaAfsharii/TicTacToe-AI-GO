package game

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
