package checkers

const (
	BOTTOM int = iota
	TOP    int = iota
)

type Rule struct {
	Rows    int
	Columns int

	// who goes first
	First byte

	// to which direction does first go
	Side int

	// how many rows should be filled
	RowsToFill int

	// should the piece become a king when it reaches the end
	BecomesKing bool

	// is multiple captures allowed in one turn
	ConsecutiveJumps bool

	// the player loses if there are no more moves left
	LoseOnNoMoves bool
}

func NewRule(rows int, columns int, first byte, side int, fill int, king bool, multiple bool, moveLoss bool) Rule {
	rule := Rule{
		Rows:             rows,
		Columns:          columns,
		First:            first,
		Side:             side,
		RowsToFill:       fill,
		BecomesKing:      king,
		ConsecutiveJumps: multiple,
		LoseOnNoMoves:    moveLoss,
	}

	return rule
}
