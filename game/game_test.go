package game

import (
	"reflect"
	"testing"

	"github.com/mraufc/tictactoe/player"
)

func TestNew(t *testing.T) {
	p1 := NewTestPlayer(nil, "p1")
	p2 := NewTestPlayer(nil, "p2")
	type args struct {
		size    int
		target  int
		player1 player.Player
		player2 player.Player
	}
	tests := []struct {
		name    string
		args    args
		want    *TicTacToe
		wantErr bool
	}{
		{
			name: "valid game 3x3, target: 3",
			args: args{
				size:    3,
				target:  3,
				player1: p1,
				player2: p2,
			},
			want: &TicTacToe{
				target: 3,
				board: [][]int{
					[]int{0, 0, 0},
					[]int{0, 0, 0},
					[]int{0, 0, 0},
				},
				player1:  p1,
				player2:  p2,
				winner:   0,
				gameOver: false,
				moves:    0,
			},
			wantErr: false,
		},
		{
			name: "valid game 6x6, target: 4",
			args: args{
				size:    6,
				target:  4,
				player1: p1,
				player2: p2,
			},
			want: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				player1:  p1,
				player2:  p2,
				winner:   0,
				gameOver: false,
				moves:    0,
			},
			wantErr: false,
		},
		{
			name: "invalid game 2x2, target: 2",
			args: args{
				size:    2,
				target:  2,
				player1: p1,
				player2: p2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid game 5x5, target: 6",
			args: args{
				size:    5,
				target:  6,
				player1: p1,
				player2: p2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid game 6x6, target: 4, nil player1",
			args: args{
				size:    6,
				target:  4,
				player1: nil,
				player2: p2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid game 6x6, target: 4, nil player2",
			args: args{
				size:    6,
				target:  4,
				player1: p1,
				player2: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.size, tt.args.target, tt.args.player1, tt.args.player2)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTicTacToe_Play(t *testing.T) {
	type want struct {
		result   bool
		gameOver bool
		winner   int
	}
	tests := []struct {
		name string
		t    *TicTacToe
		want want
	}{
		{
			name: "valid 3x3 game, target: 3, X plays 2, 2",
			t: &TicTacToe{
				target: 3,
				board: [][]int{
					[]int{1, 1, 0},
					[]int{2, 2, 0},
					[]int{0, 0, 0},
				},
				player1:  NewTestPlayer([][]int{[]int{2, 2}}, "X"),
				player2:  NewTestPlayer(nil, "O"),
				winner:   0,
				gameOver: false,
				moves:    4,
			},
			want: want{
				result:   true,
				gameOver: false,
				winner:   0,
			},
		},
		{
			name: "completed 3x3 game, target: 3",
			t: &TicTacToe{
				target: 3,
				board: [][]int{
					[]int{1, 1, 2},
					[]int{2, 2, 1},
					[]int{1, 2, 1},
				},
				player1:  NewTestPlayer([][]int{[]int{0, 2}}, "X"),
				player2:  NewTestPlayer([][]int{[]int{1, 2}}, "O"),
				winner:   0,
				gameOver: true,
				moves:    9,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   0,
			},
		},
		{
			name: "valid 3x3 game, target: 3, X plays an occupied position",
			t: &TicTacToe{
				target: 3,
				board: [][]int{
					[]int{1, 1, 0},
					[]int{2, 2, 0},
					[]int{0, 0, 0},
				},
				player1:  NewTestPlayer([][]int{[]int{0, 0}}, "X"),
				player2:  NewTestPlayer(nil, "O"),
				winner:   0,
				gameOver: false,
				moves:    4,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   2,
			},
		},
		{
			name: "valid 3x3 game, target: 3, X plays -1, 0",
			t: &TicTacToe{
				target: 3,
				board: [][]int{
					[]int{1, 1, 0},
					[]int{2, 2, 0},
					[]int{0, 0, 0},
				},
				player1:  NewTestPlayer([][]int{[]int{-1, 0}}, "X"),
				player2:  NewTestPlayer(nil, "O"),
				winner:   0,
				gameOver: false,
				moves:    4,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   2,
			},
		},
		{
			name: "valid 3x3 game, target: 3, O plays 5, 0",
			t: &TicTacToe{
				target: 3,
				board: [][]int{
					[]int{1, 1, 0},
					[]int{2, 2, 0},
					[]int{0, 0, 1},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer([][]int{[]int{5, 0}}, "O"),
				winner:   0,
				gameOver: false,
				moves:    5,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   1,
			},
		},
		{
			name: "valid 3x3 game, target: 3, X wins 0, 0 to 0, 2",
			t: &TicTacToe{
				target: 3,
				board: [][]int{
					[]int{1, 1, 0},
					[]int{2, 2, 0},
					[]int{0, 0, 0},
				},
				player1:  NewTestPlayer([][]int{[]int{0, 2}}, "X"),
				player2:  NewTestPlayer([][]int{[]int{1, 2}}, "O"),
				winner:   0,
				gameOver: false,
				moves:    4,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   1,
			},
		},
		{
			name: "valid 3x3 game, target: 3, X wins 0, 0 to 2, 2",
			t: &TicTacToe{
				target: 3,
				board: [][]int{
					[]int{1, 0, 2},
					[]int{0, 1, 2},
					[]int{0, 0, 0},
				},
				player1:  NewTestPlayer([][]int{[]int{2, 2}}, "X"),
				player2:  NewTestPlayer([][]int{[]int{2, 0}}, "O"),
				winner:   0,
				gameOver: false,
				moves:    4,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   1,
			},
		},
		{
			name: "valid 3x3 game, target: 3, X wins 0, 2 to 2, 0",
			t: &TicTacToe{
				target: 3,
				board: [][]int{
					[]int{2, 0, 1},
					[]int{0, 1, 2},
					[]int{0, 0, 0},
				},
				player1:  NewTestPlayer([][]int{[]int{2, 0}}, "X"),
				player2:  NewTestPlayer([][]int{[]int{0, 1}}, "O"),
				winner:   0,
				gameOver: false,
				moves:    4,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   1,
			},
		},
		{
			name: "valid 3x3 game, target: 3, draw",
			t: &TicTacToe{
				target: 3,
				board: [][]int{
					[]int{1, 2, 1},
					[]int{0, 1, 2},
					[]int{2, 1, 2},
				},
				player1:  NewTestPlayer([][]int{[]int{1, 0}}, "X"),
				player2:  NewTestPlayer(nil, "O"),
				winner:   0,
				gameOver: false,
				moves:    8,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   0,
			},
		},
		{
			name: "valid 6x6 game, target: 4, O wins 1, 1 to 4, 4",
			t: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{0, 1, 1, 1, 0, 1},
					[]int{0, 2, 0, 0, 0, 0},
					[]int{0, 0, 2, 0, 0, 0},
					[]int{0, 0, 0, 2, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer([][]int{[]int{4, 4}}, "O"),
				winner:   0,
				gameOver: false,
				moves:    7,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   2,
			},
		},
		{
			name: "valid 6x6 game, target: 4, O wins 0, 5 to 3, 2",
			t: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{0, 1, 1, 1, 0, 2},
					[]int{1, 0, 0, 0, 2, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 2, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer([][]int{[]int{2, 3}}, "O"),
				winner:   0,
				gameOver: false,
				moves:    7,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   2,
			},
		},
		{
			name: "valid 6x6 game, target: 4, O wins 3, 2 to 3, 5",
			t: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{0, 1, 1, 1, 0, 0},
					[]int{1, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 2, 0, 2, 2},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer([][]int{[]int{3, 3}}, "O"),
				winner:   0,
				gameOver: false,
				moves:    7,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   2,
			},
		},
		{
			name: "valid 6x6 game, target: 4, O wins 0, 0 to 0, 3",
			t: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{2, 1, 1, 1, 0, 1},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{2, 0, 0, 0, 0, 0},
					[]int{2, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer([][]int{[]int{1, 0}}, "O"),
				winner:   0,
				gameOver: false,
				moves:    7,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   2,
			},
		},
		{
			name: "valid 6x6 game, target: 4, O wins 0, 0 to 5, 5",
			t: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{2, 1, 1, 1, 0, 1},
					[]int{2, 0, 0, 0, 0, 1},
					[]int{2, 0, 0, 0, 0, 1},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{2, 0, 0, 0, 0, 0},
					[]int{2, 0, 0, 0, 0, 0},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer([][]int{[]int{3, 0}}, "O"),
				winner:   0,
				gameOver: false,
				moves:    11,
			},
			want: want{
				result:   false,
				gameOver: true,
				winner:   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Play(); got != tt.want.result {
				t.Errorf("TicTacToe.Play() = %v, want %v", got, tt.want.result)
			}
			if tt.t.gameOver != tt.want.gameOver {
				t.Errorf("TicTacToe.Play() gameOver? %v, want %v", tt.t.gameOver, tt.want.gameOver)
			}
			if tt.t.winner != tt.want.winner {
				t.Errorf("TicTacToe.Play() winner = %v, want %v", tt.t.winner, tt.want.winner)
			}
		})
	}
}

func TestTicTacToe_Evaluate(t *testing.T) {
	type args struct {
		board [][]int
		side  int
		i     int
		j     int
	}
	tests := []struct {
		name         string
		t            *TicTacToe
		args         args
		wantGameOver bool
		wantWinner   int
		wantErr      bool
	}{
		{
			name: "valid board, valid X move to 3, 3",
			t: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer(nil, "O"),
				winner:   0,
				gameOver: false,
				moves:    0,
			},
			args: args{
				board: [][]int{
					[]int{1, 0, 2, 2, 2, 0},
					[]int{0, 1, 0, 0, 0, 0},
					[]int{0, 0, 1, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				side: 1,
				i:    3,
				j:    3,
			},
			wantGameOver: true,
			wantWinner:   1,
			wantErr:      false,
		},
		{
			name: "valid board, valid X move to 0, 0",
			t: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer(nil, "O"),
				winner:   0,
				gameOver: false,
				moves:    0,
			},
			args: args{
				board: [][]int{
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				side: 1,
				i:    0,
				j:    0,
			},
			wantGameOver: false,
			wantWinner:   0,
			wantErr:      false,
		},
		{
			name: "valid board, invalid X move",
			t: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer(nil, "O"),
				winner:   0,
				gameOver: false,
				moves:    0,
			},
			args: args{
				board: [][]int{
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				side: 1,
				i:    0,
				j:    6,
			},
			wantGameOver: true,
			wantWinner:   2,
			wantErr:      false,
		},
		{
			name: "invalid board, invalid row count",
			t: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer(nil, "O"),
				winner:   0,
				gameOver: false,
				moves:    0,
			},
			args: args{
				board: [][]int{
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				side: 1,
				i:    0,
				j:    6,
			},
			wantGameOver: false,
			wantWinner:   0,
			wantErr:      true,
		},
		{
			name: "invalid board, invalid column count",
			t: &TicTacToe{
				target: 4,
				board: [][]int{
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
				},
				player1:  NewTestPlayer(nil, "X"),
				player2:  NewTestPlayer(nil, "O"),
				winner:   0,
				gameOver: false,
				moves:    0,
			},
			args: args{
				board: [][]int{
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0, 0},
					[]int{0, 0, 0, 0, 0},
				},
				side: 1,
				i:    0,
				j:    6,
			},
			wantGameOver: false,
			wantWinner:   0,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGameOver, gotWinner, err := tt.t.Evaluate(tt.args.board, tt.args.side, tt.args.i, tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("TicTacToe.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotGameOver != tt.wantGameOver {
				t.Errorf("TicTacToe.Evaluate() gotGameOver = %v, want %v", gotGameOver, tt.wantGameOver)
			}
			if gotWinner != tt.wantWinner {
				t.Errorf("TicTacToe.Evaluate() gotWinner = %v, want %v", gotWinner, tt.wantWinner)
			}
		})
	}
}

// TestPlayer implements Player interface
type TestPlayer struct {
	counter  int
	moves    [][]int
	name     string
	winner   int
	gameOver bool
}

func NewTestPlayer(moves [][]int, name string) *TestPlayer {
	return &TestPlayer{moves: moves, name: name}
}

func (tp *TestPlayer) Name() string {
	return tp.name
}

func (tp *TestPlayer) Play(board [][]int, side int) (int, int) {
	if tp.counter >= len(tp.moves) {
		return 0, 0
	}
	i, j := tp.moves[tp.counter][0], tp.moves[tp.counter][1]
	tp.counter++
	return i, j
}

func (tp *TestPlayer) Done(winner int) {
	tp.gameOver = true
	tp.winner = winner
}
