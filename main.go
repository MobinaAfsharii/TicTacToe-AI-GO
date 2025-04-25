package main

import (
	"tic-tac-toe/game"
	"tic-tac-toe/ui"
)

func main() {
	display := ui.NewDisplay()
	g := game.NewGame(display)
	g.Run()
}
