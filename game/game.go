package game

import (
	"fmt"
	"tic-tac-toe/ai"
	"tic-tac-toe/ui"
)


type Displayer interface {
	ClearScreen()
	RenderBoard(board Board, playerMarker, aiMarker Marker, winningCombo WinningCombo)
	RenderScores(scores Score)
	PromptMarker() (Marker, error)
	PromptFirstTurn() (PlayerType, error)
	PromptMove(board Board, marker Marker) (int, error)
	PromptPlayAgain() (bool, error)
	ShowMessage(msg string, style ui.MessageStyle)
}


type Game struct {
	display      Displayer
	board        Board
	playerMarker Marker
	aiMarker     Marker
	currentTurn  PlayerType
	scores       Score
}


func NewGame(display Displayer) *Game {
	return &Game{
		display: display,
		scores:  Score{},
	}
}


func (g *Game) Run() {
	for {
		g.setupGame()
		g.playGame()
		if !g.playAgain() {
			g.display.ShowMessage("Thanks for playing!", ui.InfoMessage)
			break
		}
	}
}


func (g *Game) setupGame() {
	g.board = NewBoard()
	var err error
	g.playerMarker, err = g.display.PromptMarker()
	if err != nil {
		g.display.ShowMessage("Invalid marker, defaulting to X", ui.ErrorMessage)
		g.playerMarker = MarkerX
	}
	g.aiMarker, _ = g.playerMarker.OppositeMarker()
	g.currentTurn, err = g.display.PromptFirstTurn()
	if err != nil {
		g.display.ShowMessage("Invalid choice, player goes first", ui.ErrorMessage)
		g.currentTurn = PlayerHuman
	}
}


func (g *Game) playGame() {
	for {
		g.display.ClearScreen()
		g.display.RenderBoard(g.board, g.playerMarker, g.aiMarker, WinningCombo{})
		g.display.RenderScores(g.scores)

		if g.currentTurn == PlayerHuman {
			g.handlePlayerMove()
		} else {
			g.handleAIMove()
		}

		if g.isGameOver() {
			g.display.RenderBoard(g.board, g.playerMarker, g.aiMarker, g.getWinningCombo())
			g.display.RenderScores(g.scores)
			break
		}

		g.switchTurn()
	}
}


func (g *Game) handlePlayerMove() {
	for {
		move, err := g.display.PromptMove(g.board, g.playerMarker)
		if err == nil && g.board.MakeMove(move, g.playerMarker) == nil {
			break
		}
		g.display.ShowMessage("Invalid move, try again", ui.ErrorMessage)
	}
}


func (g *Game) handleAIMove() {
	g.display.ShowMessage(fmt.Sprintf("AI's turn (%s)...", g.aiMarker), ui.InfoMessage)
	move := ai.FindBestMove(g.board, g.playerMarker, g.aiMarker)
	if move == -1 || g.board.MakeMove(move, g.aiMarker) != nil {
		g.display.ShowMessage("AI failed to make a move", ui.ErrorMessage)
	}
}


func (g *Game) switchTurn() {
	if g.currentTurn == PlayerHuman {
		g.currentTurn = PlayerAI
	} else {
		g.currentTurn = PlayerHuman
	}
}


func (g *Game) isGameOver() bool {
	if combo, won := g.board.CheckWin(g.playerMarker); won {
		g.scores.Player++
		g.display.ShowMessage("🎉 You win! 🎉", ui.WinMessage)
		return true
	}
	if combo, won := g.board.CheckWin(g.aiMarker); won {
		g.scores.AI++
		g.display.ShowMessage("🤖 AI wins! 🤖", ui.ErrorMessage)
		return true
	}
	if g.board.IsFull() {
		g.scores.Draws++
		g.display.ShowMessage("🤝 It's a draw! 🤝", ui.DrawMessage)
		return true
	}
	return false
}


func (g *Game) getWinningCombo() WinningCombo {
	if combo, won := g.board.CheckWin(g.playerMarker); won {
		return combo
	}
	if combo, won := g.board.CheckWin(g.aiMarker); won {
		return combo
	}
	return WinningCombo{}
}


func (g *Game) playAgain() bool {
	playAgain, err := g.display.PromptPlayAgain()
	if err != nil {
		return false
	}
	return playAgain
}
