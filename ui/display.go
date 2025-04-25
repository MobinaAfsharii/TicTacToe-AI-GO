package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tic-tac-toe/game"

	"github.com/charmbracelet/lipgloss"
)


type MessageStyle string

const (
	WinMessage   MessageStyle = "win"
	DrawMessage  MessageStyle = "draw"
	ErrorMessage MessageStyle = "error"
	InfoMessage  MessageStyle = "info"
)


type Display struct {
	scanner    *bufio.Scanner
	styles     map[string]lipgloss.Style
	boardStyle lipgloss.Style
}


func NewDisplay() *Display {
	return &Display{
		scanner: bufio.NewScanner(os.Stdin),
		styles: map[string]lipgloss.Style{
			"title": lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("201")).Padding(1, 2).BorderStyle(lipgloss.RoundedBorder()),
			"player": lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Bold(true),
			"ai":     lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true),
			"win":    lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true).Padding(1, 2).BorderStyle(lipgloss.RoundedBorder()),
			"draw":   lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Bold(true).Padding(1, 2).BorderStyle(lipgloss.RoundedBorder()),
			"error":  lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true).Padding(1, 2),
			"info":   lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Padding(1, 2),
			"score":  lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Padding(1, 2).BorderStyle(lipgloss.RoundedBorder()),
		},
		boardStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("7")),
	}
}


func (d *Display) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}


func (d *Display) RenderBoard(board game.Board, playerMarker, aiMarker game.Marker, winningCombo game.WinningCombo) {
	d.ClearScreen()
	fmt.Println(d.styles["title"].Render("Tic-Tac-Toe AI"))

	grid := "\n"
	for i := 0; i < 9; i += 3 {
		row := make([]string, 3)
		for j := 0; j < 3; j++ {
			idx := i + j
			content := string(board[idx])
			style := d.boardStyle
			if board[idx] == playerMarker {
				style = d.styles["player"]
			} else if board[idx] == aiMarker {
				style = d.styles["ai"]
			} else {
				content = fmt.Sprintf("%d", idx+1)
			}
			if winningCombo != [3]int{} && (idx == winningCombo[0] || idx == winningCombo[1] || idx == winningCombo[2]) {
				style = d.styles["win"]
			}
			row[j] = style.Render(fmt.Sprintf(" %s ", content))
		}
		grid += d.boardStyle.Render(fmt.Sprintf(" %s | %s | %s \n", row[0], row[1], row[2]))
		if i < 6 {
			grid += d.boardStyle.Render("---+---+---\n")
		}
	}
	fmt.Println(grid)
}


func (d *Display) RenderScores(scores game.Score) {
	scoreText := fmt.Sprintf("Player: %d | AI: %d | Draws: %d", scores.Player, scores.AI, scores.Draws)
	fmt.Println(d.styles["score"].Render(scoreText))
}


func (d *Display) PromptMarker() (game.Marker, error) {
	fmt.Printf("Choose your marker (%s or %s): ", game.MarkerX, game.MarkerO)
	d.scanner.Scan()
	input := strings.TrimSpace(strings.ToUpper(d.scanner.Text()))
	marker := game.Marker(input)
	if marker.IsValid() {
		return marker, nil
	}
	return game.Empty, fmt.Errorf("invalid marker: %s", input)
}


func (d *Display) PromptFirstTurn() (game.PlayerType, error) {
	fmt.Print("Who goes first? (Player or AI): ")
	d.scanner.Scan()
	input := strings.TrimSpace(strings.Title(d.scanner.Text()))
	playerType := game.PlayerType(input)
	for _, pt := range game.ValidPlayerTypes() {
		if pt == playerType {
			return playerType, nil
		}
	}
	return "", fmt.Errorf("invalid player type: %s", input)
}


func (d *Display) PromptMove(board game.Board, marker game.Marker) (int, error) {
	fmt.Printf("Your turn (%s). Enter move (1-9): ", marker)
	d.scanner.Scan()
	input := strings.TrimSpace(d.scanner.Text())
	move, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", input)
	}
	if move < 1 || move > 9 {
		return 0, fmt.Errorf("move out of range: %d", move)
	}
	if board[move-1] != game.Empty {
		return 0, fmt.Errorf("cell %d is not empty", move)
	}
	return move - 1, nil
}


func (d *Display) PromptPlayAgain() (bool, error) {
	fmt.Print("Play again? (y/n): ")
	d.scanner.Scan()
	input := strings.TrimSpace(strings.ToLower(d.scanner.Text()))
	switch input {
	case "y":
		return true, nil
	case "n":
		return false, nil
	default:
		return false, fmt.Errorf("invalid input: %s", input)
	}
}


func (d *Display) ShowMessage(msg string, style MessageStyle) {
	fmt.Println(d.styles[string(style)].Render(msg))
}
