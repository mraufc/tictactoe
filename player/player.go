// Package player defines the Player interface that is used by the TicTacToe game.
package player

// Player is the basic TicTacToe variation player
type Player interface {
	// Play returns a position to play given the board and player side.
	Play(board [][]int, side int) (int, int)
	// Done informs the player that the current game is over.
	// winner 0 means the game is a tie, 1 means player 1 won and 2 means player 2 won.
	Done(winner int)
	// Name returns the player name / id.
	Name() string
}
