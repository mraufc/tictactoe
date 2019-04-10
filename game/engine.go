package game

// Engine is the game engine that evaluates moves for a board of size rows by columns.
type Engine struct {
	target  int
	rows    int
	columns int
}

// NewEngine returns a new game engine
func NewEngine(rows, columns, target int) (*Engine, error) {
	if rows < 3 || columns < 3 || target < 3 || target > rows || target > columns {
		return nil, ErrInvalidGameSpecs
	}
	return &Engine{
		target:  target,
		rows:    rows,
		columns: columns,
	}, nil
}

// Evaluate evalutes a hypothetical board position and a side's move.
// board parameter and the TicTacToe's instance board sizes must match.
// Side is 1 for X Player and 2 for O Player.
// Evaluate function returns whether the game will be over after the move, and the winner of the game
// if the game is over. Winner can be 0 for draw, 1 for X Player and 2 for O Player.
// TODO: maybe "side" and "winner" can be enums.
func (e *Engine) Evaluate(board [][]int, side, i, j int) (gameOver bool, winner int, err error) {
	if board == nil {
		err = ErrInvalidBoard
		return
	}
	if side != 1 && side != 2 {
		err = ErrInvalidSide
		return
	}
	if len(board) != e.rows {
		err = ErrInvalidBoard
		return
	}
	unoccupied := 0
	for _, row := range board {
		if len(row) != e.columns {
			err = ErrInvalidBoard
			return
		}
		for _, v := range row {
			if v == 0 {
				unoccupied++
			}
		}
	}
	gameOver, winner = e.evaluate(board, side, i, j, unoccupied)
	return
}

func (e *Engine) evaluate(board [][]int, side, i, j, unoccupied int) (bool, int) {
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
	if i-e.target+1 > minI {
		minI = i - e.target + 1
	}
	if i+e.target-1 < maxI {
		maxI = i + e.target - 1
	}

	minJ, maxJ := 0, len(board)-1
	if j-e.target+1 > minJ {
		minJ = j - e.target + 1
	}
	if j+e.target-1 < maxJ {
		maxJ = j + e.target - 1
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
	if cnt >= e.target {
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
	if cnt >= e.target {
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
	if cnt >= e.target {
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
	if cnt >= e.target {
		return true, side
	}
	if unoccupied == 1 {
		return true, 0
	}
	return false, 0
}
