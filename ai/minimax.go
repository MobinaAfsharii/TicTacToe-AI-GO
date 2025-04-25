package ai

import (
	"math"
	"math/rand"
	"tic-tac-toe/types"
	"time"
)

var positionalScores = [9]float64{
	0.3, 0.1, 0.3,
	0.1, 0.5, 0.1,
	0.3, 0.1, 0.3,
}

func init() {

	rand.Seed(time.Now().UnixNano())
}

func FindBestMove(board types.Board, playerMarker, aiMarker types.Marker) int {

	for _, move := range board.AvailableMoves() {

		board.MakeMove(move, aiMarker)
		if _, won := board.CheckWin(aiMarker); won {
			board.MakeMove(move, types.Empty)
			return move
		}
		board.MakeMove(move, types.Empty)

		board.MakeMove(move, playerMarker)
		if _, won := board.CheckWin(playerMarker); won {
			board.MakeMove(move, types.Empty)
			return move
		}
		board.MakeMove(move, types.Empty)
	}

	bestMoves := []int{}
	bestScore := math.Inf(-1)
	alpha := math.Inf(-1)
	beta := math.Inf(1)
	availableMoves := board.AvailableMoves()

	for _, move := range availableMoves {
		board.MakeMove(move, aiMarker)
		score := minimax(board, 0, false, playerMarker, aiMarker, alpha, beta)

		score += positionalScores[move] * 0.01
		board.MakeMove(move, types.Empty)

		if score > bestScore {
			bestScore = score
			bestMoves = []int{move}
		} else if score == bestScore {
			bestMoves = append(bestMoves, move)
		}
		alpha = math.Max(alpha, bestScore)
	}

	if len(bestMoves) == 0 {
		return -1
	}

	return bestMoves[rand.Intn(len(bestMoves))]
}

func minimax(board types.Board, depth int, isMaximizing bool, playerMarker, aiMarker types.Marker, alpha, beta float64) float64 {
	if _, won := board.CheckWin(aiMarker); won {
		return float64(10 - depth)
	}
	if _, won := board.CheckWin(playerMarker); won {
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
			score := minimax(board, depth+1, false, playerMarker, aiMarker, alpha, beta)
			board.MakeMove(move, types.Empty)
			bestScore = math.Max(bestScore, score)
			alpha = math.Max(alpha, bestScore)
			if beta <= alpha {
				break
			}
		}
		return bestScore
	} else {
		bestScore := math.Inf(1)
		for _, move := range availableMoves {
			board.MakeMove(move, playerMarker)
			score := minimax(board, depth+1, true, playerMarker, aiMarker, alpha, beta)
			board.MakeMove(move, types.Empty)
			bestScore = math.Min(bestScore, score)
			beta = math.Min(beta, bestScore)
			if beta <= alpha {
				break
			}
		}
		return bestScore
	}
}
