// Package game is a simple generalized TicTacToe variant implementation.
// The implementation allows a flexible N x M game size where N >= 3, M >= 3 and T consequetive symbols
// as win condition where T >= 3, T <= N and T <= M.
package game

import (
	"fmt"

	"github.com/mraufc/tictactoe/player"
)

// TicTacToe is a an NxN board game where two sides (X and O) take turns to place their symbols.
// First player to reach a certain number (indicated by Engine's target) of X's or O's vertically, horizontally
// or diagonally wins the game.
type TicTacToe struct {
	board    [][]int // 0 -> empty position, 1 -> X, 2 -> O
	player1  player.Player
	player2  player.Player
	winner   int // 0 is a tie or draw
	gameOver bool
	moves    int
	e        *Engine
}

// New returns a new game of TicTacToe.
// Board size must be greater than or equal to 3.
// Win condition (target) must between 3 and board size.
func New(engine *Engine, player1, player2 player.Player) (*TicTacToe, error) {
	if engine == nil || player1 == nil || player2 == nil {
		return nil, ErrInvalidGameSpecs
	}
	board := make([][]int, engine.rows)
	for i := range board {
		board[i] = make([]int, engine.columns)
	}
	return &TicTacToe{
		board:   board,
		e:       engine,
		player1: player1,
		player2: player2,
	}, nil
}

// Play calls the Play function of the appropriate player and evaluates the move and board.
// This function returns true as long as game is not over.
func (t *TicTacToe) Play() bool {
	if t.gameOver {
		return false
	}

	// pass a copy of the board to the player
	cpy := make([][]int, len(t.board))
	for i, row := range t.board {
		cpy[i] = make([]int, len(row))
		for j, v := range row {
			cpy[i][j] = v
		}
	}
	var i, j, side int
	if t.moves%2 == 0 {
		side = 1
		i, j = t.player1.Play(cpy, 1)
	} else {
		side = 2
		i, j = t.player2.Play(cpy, 2)
	}
	t.gameOver, t.winner = t.e.evaluate(t.board, side, i, j, len(t.board)*len(t.board)-t.moves)
	if t.gameOver {
		t.player1.Done(t.winner)
		t.player2.Done(t.winner)
		// illegal move, do not update the board
		if t.winner != side && t.winner != 0 {
			return false
		}
	}
	t.board[i][j] = side
	t.moves++
	return !t.gameOver
}

// Result returns if the game is still in progress and the winner
func (t *TicTacToe) Result() (bool, int) {
	return !t.gameOver, t.winner
}

// Pretty returns a pretty string representation of the board
func (t *TicTacToe) Pretty() string {
	size := len(t.board)
	title := fmt.Sprintf("%v as 'X' vs. %v as 'O'\n", t.player1.Name(), t.player2.Name())
	board := ""
	for i := 0; i < size; i++ {
		line := ""
		for j := 0; j < size; j++ {
			if t.board[i][j] == 0 {
				line += "-"
			} else if t.board[i][j] == 1 {
				line += "X"
			} else if t.board[i][j] == 2 {
				line += "O"
			}
			if j < size-1 {
				line += " "
			}
		}
		line += "\n"
		board += line
	}
	var result string
	if !t.gameOver {
		result = "Game is still in progress"
	} else {
		switch t.winner {
		case 1:
			result = fmt.Sprintf("Winner is %v as 'X'", t.player1.Name())
		case 2:
			result = fmt.Sprintf("Winner is %v as 'O'", t.player2.Name())
		case 0:
			result = "Game is a Draw!"
		}
	}
	return title + board + result
}
