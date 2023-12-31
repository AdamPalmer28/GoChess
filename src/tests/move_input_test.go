package test

import (
	"chess/src/chess_engine/board"
	"testing"
)

func TestCordIndex(t *testing.T) {

	test_cases := []struct {
		cord     string
		expected uint
	}{
		{"a1", 0},
		{"a3", 16},
		{"d4", 27},
		{"g1", 6},
		{"h8", 63},
		{"e5", 36},
		{"b7", 49},
	}

	for _, tc := range test_cases {

		result := board.Move_to_index(tc.cord)
		if result != tc.expected {
			t.Errorf("move_to_index(%s) = %d; want %d", tc.cord, result, tc.expected)
		}
		cord_result := board.Index_to_move(tc.expected)
		if cord_result != tc.cord {
			t.Errorf("index_to_move(%d) = %s; want %s", tc.expected, cord_result, tc.cord)
		}
	}
}
