package ai

import (
	"math"
	"tic-tac-toe/game"
)


func FindBestMove(board game.Board, playerMarker, aiMarker game.Marker) int {
	bestScore := math.Inf(-1)
	bestMove := -1
	availableMoves := board.AvailableMoves()

	for _, move := range availableMoves {
		board.MakeMove(move, aiMarker)
		score := minimax(board, 0, false, playerMarker, aiMarker)
		board.MakeMove(move, game.Empty) 
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	if bestMove == -1 && len(availableMoves) > 0 {
		return availableMoves[0] 
	}
	return bestMove
}


func minimax(board game.Board, depth int, isMaximizing bool, playerMarker, aiMarker game.Marker) float64 {
	if combo, won := board.CheckWin(aiMarker); won {
		return float64(10 - depth)
	}
	if combo, won := board.CheckWin(playerMarker); won {
		return float64(depth - 10)
	}
	if board.IsFull() {
		return 0
	}

	availableMoves := board.AvailableMoves()
	if isMaximizing {
		bestScore := math.Inf(-1)
		for _, move := range availableMoves {
			board.MakeMove(move, aiMarker)
			score := minimax(board, depth+1, false, playerMarker, aiMarker)
			board.MakeMove(move, game.Empty)
			bestScore = math.Max(bestScore, score)
		}
		return bestScore
	} else {
		bestScore := math.Inf(1)
		for _, move := range availableMoves {
			board.MakeMove(move, playerMarker)
			score := minimax(board, depth+1, true, playerMarker, aiMarker)
			board.MakeMove(move, game.Empty)
			bestScore = math.Min(bestScore, score)
		}
		return bestScore
	}
}
