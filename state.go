package checkers


type State struct {
	rule Rule

	board [][]uint
}


func NewState(rule Rule) *State {
	state := State{
		rule: rule
		board: make([][]uint, rows)
	}

	for i := 0; i < rows; i++ {
		state.board[i] = make([]uint, columns)
	}

	return &state
}

func (state *State) ValidMove(coord Coordinate) bool {
}
