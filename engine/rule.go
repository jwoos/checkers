package checkers

// TODO change to directional values -1 and 1
const (
	BOTTOM byte = iota
	TOP    byte = iota
)

type Rule struct {
	Rows    int
	Columns int

	// who goes first
	First byte

	// to which direction does first go
	Side byte

	Order []byte
	Direction []byte

	// how many rows should be filled
	RowsToFill int

	// should the piece become a king when it reaches the end
	BecomesKing bool

	// is multiple captures allowed in one turn
	ConsecutiveJumps bool

	// the player loses if there are no more moves left
	LoseOnNoMoves bool
}

func NewRule(rows int, columns int, first byte, side byte, fill int, king bool, multiple bool, moveLoss bool) Rule {
	rule := Rule{
		Rows:             rows,
		Columns:          columns,
		Order: make([]byte, 2),
		Direction: make([]byte, 3),
		First:            first,
		Side:             side,
		RowsToFill:       fill,
		BecomesKing:      king,
		ConsecutiveJumps: multiple,
		LoseOnNoMoves:    moveLoss,
	}

	if first == BLACK {
		rule.Direction[BLACK] = side
		rule.Direction[WHITE] = side ^ TOP

		rule.Order[0] = BLACK
		rule.Order[1] = WHITE
	} else {
		rule.Direction[WHITE] = side
		rule.Direction[BLACK] = side ^ TOP

		rule.Order[0] = BLACK
		rule.Order[1] = WHITE
	}

	return rule
}
