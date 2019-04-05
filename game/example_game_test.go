package game_test

import (
	"fmt"

	"github.com/mraufc/tictactoe/game"
)

func ExampleTicTacToe_Pretty() {
	p1 := game.NewTestPlayer([][]int{[]int{0, 0}, []int{0, 1}, []int{0, 2}, []int{0, 4}}, "Player 1")
	p2 := game.NewTestPlayer([][]int{[]int{1, 0}, []int{1, 1}, []int{1, 2}, []int{1, 3}}, "Player 2")
	t, err := game.New(6, 4, p1, p2)
	if err != nil {
		fmt.Println(err)
		return
	}

	inProgress := true
	for inProgress {
		inProgress = t.Play()
	}
	fmt.Print(t.Pretty())

	// Output:
	// Player 1 as 'X' vs. Player 2 as 'O'
	// X X X - X -
	// O O O O - -
	// - - - - - -
	// - - - - - -
	// - - - - - -
	// - - - - - -
	// Winner is Player 2 as 'O'
}
