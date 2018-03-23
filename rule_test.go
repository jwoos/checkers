package checkers

import (
	"reflect"
	"testing"
)

func TestNewRule(t *testing.T) {
	rule := NewRule(10, 10, BLACK, TOP, 3, false, false, false)
	expected := Rule{
		Rows:             10,
		Columns:          10,
		First:            BLACK,
		Side:             TOP,
		Order: make([]byte, 2),
		Direction: make([]byte, 3),
		RowsToFill:       3,
		BecomesKing:      false,
		ConsecutiveJumps: false,
		LoseOnNoMoves:    false,
	}

	expected.Order[0] = BLACK
	expected.Order[1] = WHITE

	expected.Direction[BLACK] = TOP
	expected.Direction[WHITE] = BOTTOM

	if !reflect.DeepEqual(rule, expected) {
		t.Errorf("Expected %v and got %v", expected, rule)
	}
}
