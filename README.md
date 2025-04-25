# Tic-Tac-Toe AI (Go)

A console-based Tic-Tac-Toe game with an unbeatable AI opponent, implemented in Go using the Minimax algorithm.

## Features
- Play against an AI that makes optimal moves using the Minimax algorithm.
- Clean console UI with colored output using the `lipgloss` library.
- Choose your marker (X or O) and who goes first.
- Tracks scores across multiple games and supports replay.
- Highlights winning combinations and game outcomes (win, draw).

## Prerequisites
- Go 1.18 or later
- `lipgloss` library: `go get github.com/charmbracelet/lipgloss`

## Installation
1. Clone the repository or copy the code into a directory.
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the game:
   ```bash
   go run .
   ```

## Directory Structure
```
tic-tac-toe/
├── game/
│   ├── board.go        # Board logic and game state
│   ├── game.go         # Game loop and mechanics
│   └── types.go        # Shared types and constants
├── ai/
│   └── minimax.go      # AI logic with Minimax algorithm
├── ui/
│   └── display.go      # Console UI rendering
├── go.mod              # Go module file
└── main.go             # Entry point
```

## How to Play
1. Run the game with `go run .`.
2. Choose your marker (X or O) and who goes first (Player or AI).
3. Enter moves by selecting a cell (1-9, corresponding to the board positions).
4. The AI will respond with optimal moves.
5. After a game ends (win, draw), choose to play again or exit.
6. Scores are displayed after each game.

## Gameplay
- The board is a 3x3 grid, numbered 1-9 (left to right, top to bottom).
- Players alternate placing their marker (X or O) in an empty cell.
- The game ends with a win (three markers in a row, column, or diagonal) or a draw (board full).
- The AI uses the Minimax algorithm, ensuring it never loses.

## Best Practices
- **Modularity**: Logic is split into `game`, `ai`, and `ui` packages.
- **Type Safety**: Custom types for board, markers, and game states.
- **Error Handling**: Robust input validation and error propagation.
- **Idiomatic Go**: Clear function names, minimal global state, and explicit returns.
- **Minimal Dependencies**: Only `lipgloss` for UI styling.

## License
MIT License
