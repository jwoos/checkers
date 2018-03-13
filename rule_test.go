package checkers

import (
	"testing"
)

func TestNewRule(t *testing.T) {
	rule := NewRule(10, 10, BLACK, TOP, 3, false, false, false)
	expected := Rule{
		Rows:             10,
		Columns:          10,
		First:            BLACK,
		Side:             TOP,
		RowsToFill:       3,
		BecomesKing:      false,
		ConsecutiveJumps: false,
		LoseOnNoMoves:    false,
	}

	if rule != expected {
		t.Fail()
	}
}
