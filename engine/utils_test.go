package checkers

import (
	"testing"
)

func TestMin(t *testing.T) {
	cases := [][]int{
		{0, 1, 0},
		{1, 2, 1},
		{-10, 5, -10},
		{100, 101, 100},
		{-10, -15, -15},
	}

	for _, v := range cases {
		if min(v[0], v[1]) != v[2] {
			t.Fail()
		}
	}
}

func TestMax(t *testing.T) {
	cases := [][]int{
		{0, 1, 1},
		{1, 2, 2},
		{-10, 5, 5},
		{100, 101, 101},
		{-10, -15, -10},
	}

	for _, v := range cases {
		if max(v[0], v[1]) != v[2] {
			t.Fail()
		}
	}
}
