package checkers


type Rule struct {
	Rows uint
	Columns uint

	// should the piece become a king when it reaches the end
	BecomesKing bool

	// is multiple captures allowed in one turn
	ConsecutiveJumps bool

	// the player loses if there are no more moves left
	LoseOnNoMoves bool
}


func NewRule(rows uint, columns uint, king bool, multiple bool, moveLoss bool) Rule {
	rule := Rules{
		Rows: rows,
		Columns: columns,
		BecomesKing: king,
		ConsecutiveJumps: multiple,
		LoseOnNoMoves: moveLoss
	}

	return rule
}
