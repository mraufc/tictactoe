// Package game is a simple TicTacToe variant implementation.
// The implementation allows a flexible N x N game size where N >= 3 and M consequetive symbols
// as win condition where M >= 3 & M <= N.
package game

import (
	"fmt"

	"github.com/mraufc/tictactoe/player"
)

// TicTacToe is a an NxN board game where two sides (X and O) take turns to place their symbols.
// First player to reach a certain number (indicated by target) of X's or O's vertically, horizontally
// or diagonally wins the game.
type TicTacToe struct {
	target   int
	board    [][]int // 0 -> empty position, 1 -> X, 2 -> O
	player1  player.Player
	player2  player.Player
	winner   int // 0 is a tie or draw
	gameOver bool
	moves    int
}

// New returns a new game of TicTacToe.
// Board size must be greater than or equal to 3.
// Win condition (target) must between 3 and board size.
func New(size, target int, player1, player2 player.Player) (*TicTacToe, error) {
	if size < 3 || target < 3 || target > size || player1 == nil || player2 == nil {
		return nil, ErrInvalidGameSpecs
	}
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}
	return &TicTacToe{
		board:   board,
		target:  target,
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
	t.gameOver, t.winner = t.evaluate(t.board, side, i, j, len(t.board)*len(t.board)-t.moves)
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

// Evaluate evalutes a hypothetical board position and a side's move.
// board parameter and the TicTacToe's instance board sizes must match.
func (t *TicTacToe) Evaluate(board [][]int, side, i, j int) (gameOver bool, winner int, err error) {
	if board == nil {
		err = ErrInvalidBoard
		return
	}
	if side != 1 && side != 2 {
		err = ErrInvalidSide
		return
	}
	if len(board) != len(t.board) {
		err = ErrInvalidBoard
		return
	}
	unoccupied := 0
	for _, row := range board {
		if len(row) != len(t.board) {
			err = ErrInvalidBoard
			return
		}
		for _, v := range row {
			if v == 0 {
				unoccupied++
			}
		}
	}
	gameOver, winner = t.evaluate(board, side, i, j, unoccupied)
	return
}

func (t *TicTacToe) evaluate(board [][]int, side, i, j, unoccupied int) (bool, int) {
	// if there are no unoccupied positions left, the game is already over
	if unoccupied == 0 {
		return true, 0
	}
	// if the player makes an invalid move or the move position is already occupied,
	// that player loses immediately.
	if i < 0 || j < 0 || i >= len(board) || j >= len(board) || board[i][j] != 0 {
		if side == 1 {
			return true, 2 // winner is 2 (O)
		}
		return true, 1 // winner is 1 (X)
	}

	minI, maxI := 0, len(board)-1
	if i-t.target+1 > minI {
		minI = i - t.target + 1
	}
	if i+t.target-1 < maxI {
		maxI = i + t.target - 1
	}

	minJ, maxJ := 0, len(board)-1
	if j-t.target+1 > minJ {
		minJ = j - t.target + 1
	}
	if j+t.target-1 < maxJ {
		maxJ = j + t.target - 1
	}

	// check vertical
	cnt := 1
	for k := 1; i-k >= minI; k++ {
		if board[i-k][j] != side {
			break
		}
		cnt++
	}
	for k := 1; i+k <= maxI; k++ {
		if board[i+k][j] != side {
			break
		}
		cnt++
	}
	if cnt >= t.target {
		return true, side
	}

	// check horizontal
	cnt = 1
	for k := 1; j-k >= minJ; k++ {
		if board[i][j-k] != side {
			break
		}
		cnt++
	}
	for k := 1; j+k <= maxJ; k++ {
		if board[i][j+k] != side {
			break
		}
		cnt++
	}
	if cnt >= t.target {
		return true, side
	}

	// check diagonal upper left to lower right
	cnt = 1
	for k := 1; i-k >= minI && j-k >= minJ; k++ {
		if board[i-k][j-k] != side {
			break
		}
		cnt++
	}
	for k := 1; i+k <= maxI && j+k <= maxJ; k++ {
		if board[i+k][j+k] != side {
			break
		}
		cnt++
	}
	if cnt >= t.target {
		return true, side
	}

	// check diagonal upper right to lower left
	cnt = 1
	for k := 1; i-k >= minI && j+k <= maxJ; k++ {
		if board[i-k][j+k] != side {
			break
		}
		cnt++
	}
	for k := 1; i+k <= maxI && j-k >= minJ; k++ {
		if board[i+k][j-k] != side {
			break
		}
		cnt++
	}
	if cnt >= t.target {
		return true, side
	}
	if unoccupied == 1 {
		return true, 0
	}
	return false, 0
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
