package main

import "testing"

// BFFFBBFRRR: row 70, column 7, seat ID 567.
// FFFBBBFRRR: row 14, column 7, seat ID 119.
// BBFFBBFRLL: row 102, column 4, seat ID 820.

func TestGetSeatID(t *testing.T) {

	type testcase struct {
		Name           string
		Instructions   string
		ExpectedRow    int
		ExpectedColumn int
		ExpectedSeatID int
	}

	testcases := []testcase{
		testcase{
			Name:           "first",
			Instructions:   "BFFFBBFRRR",
			ExpectedRow:    70,
			ExpectedColumn: 7,
			ExpectedSeatID: 567,
		},
		testcase{
			Name:           "second",
			Instructions:   "FFFBBBFRRR",
			ExpectedRow:    14,
			ExpectedColumn: 7,
			ExpectedSeatID: 119,
		},
		testcase{
			Name:           "third",
			Instructions:   "BBFFBBFRLL",
			ExpectedRow:    102,
			ExpectedColumn: 4,
			ExpectedSeatID: 820,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {

			y, x := getSeatPosition(tc.Instructions)

			if y != tc.ExpectedRow {
				t.Fatalf("expected row: %d, got: %d", tc.ExpectedRow, y)
			}

			if x != tc.ExpectedColumn {
				t.Fatalf("expected column: %d, got: %d", tc.ExpectedColumn, x)
			}

			if seatID := getSeatID(y, x); seatID != tc.ExpectedSeatID {
				t.Fatalf("expected seatID: %d, got: %d", tc.ExpectedSeatID, seatID)
			}
		})
	}
}
