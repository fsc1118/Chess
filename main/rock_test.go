package main

import (
	"strconv"
	"testing"
)

func TestRockValidMove(t *testing.T) {
	chessBoard = newBoard()
	rock1 := rock{pieceEntity{WHITE, 0, 0}}
	res1 := rock1.isMoveValid(0, 1)
	if res1 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s.", strconv.FormatBool(res1), strconv.FormatBool(true))
	}
	rock2 := rock{pieceEntity{WHITE, 0, 0}}
	res2 := rock2.isMoveValid(1, 0)
	if res2 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res2), strconv.FormatBool(true))
	}
	rock3 := rock{pieceEntity{WHITE, 0, 0}}
	res3 := rock3.isMoveValid(1, 1)
	if res3 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res3), strconv.FormatBool(false))
	}
	rock4 := rock{pieceEntity{WHITE, 0, 0}}
	res4 := rock4.isMoveValid(0, 0)
	if res4 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res4), strconv.FormatBool(false))
	}
	rock5 := rock{pieceEntity{WHITE, 0, 0}}
	res5 := rock5.isMoveValid(0, 7)
	if res5 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res5), strconv.FormatBool(true))
	}
	rock6 := rock{pieceEntity{WHITE, 0, 0}}
	res6 := rock6.isMoveValid(7, 0)
	if res6 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res6), strconv.FormatBool(true))
	}
	rock7 := rock{pieceEntity{WHITE, 0, 0}}
	res7 := rock7.isMoveValid(7, 7)
	if res7 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res7), strconv.FormatBool(false))
	}
	rock8 := rock{pieceEntity{WHITE, 0, 0}}
	// check if the rock can move to a position where it is blocked by another piece
	chessBoard.board[0][1].isEmpty = false
	res8 := rock8.isMoveValid(0, 7)
	if res8 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res7), strconv.FormatBool(false))
	}
}

func TestRockDisplay(t *testing.T) {
	chessBoard = newBoard()
	rock1 := rock{pieceEntity{WHITE, 0, 0}}
	res1 := rock1.display()
	if res1 != "WR" {
		t.Errorf("Display was incorrect, got: %s, want: %s", res1, "WR")
	}
	rock2 := rock{pieceEntity{BLACK, 0, 0}}
	res2 := rock2.display()
	if res2 != "BR" {
		t.Errorf("Display was incorrect, got: %s, want: %s.", res2, "BR")
	}
}
