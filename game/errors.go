package game

import "errors"

// ErrInvalidBoard is returned when board input is not valid
var ErrInvalidBoard = errors.New("invalid board")

// ErrInvalidGameSpecs is returned when board size and/or game win condition numbers are invalid
// or player1 or player2 is nil.
var ErrInvalidGameSpecs = errors.New("invalid game specifications")

// ErrInvalidSide is returned when side is invalid
var ErrInvalidSide = errors.New("invalid side")
