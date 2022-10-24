package main

import (
	"strconv"
	"testing"
)

func TestPawnMoveValid(t *testing.T) {
	chessBoard = newBoard()
	pawn1 := pawn{pieceEntity{BLACK, 0, 0}, false}
	res1 := pawn1.isMoveValid(0, -1)
	if res1 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s.", strconv.FormatBool(res1), strconv.FormatBool(false))
	}
	pawn2 := pawn{pieceEntity{WHITE, 0, 0}, false}
	res2 := pawn2.isMoveValid(2, 0)
	if res2 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res2), strconv.FormatBool(false))
	}
	pawn3 := pawn{pieceEntity{WHITE, 0, 0}, true}
	res3 := pawn3.isMoveValid(2, 0)
	if res3 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res3), strconv.FormatBool(true))
	}
	pawn4 := pawn{pieceEntity{WHITE, 0, 0}, false}
	res4 := pawn4.isMoveValid(1, 0)
	if res4 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res4), strconv.FormatBool(true))
	}
	pawn5 := pawn{pieceEntity{WHITE, 0, 0}, false}
	res5 := pawn5.isMoveValid(1, 2)
	if res5 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res5), strconv.FormatBool(false))
	}
	pawn6 := pawn{pieceEntity{WHITE, 0, 0}, false}
	res6 := pawn6.isMoveValid(2, 1)
	if res6 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res6), strconv.FormatBool(false))
	}
	pawn7 := pawn{pieceEntity{WHITE, 1, 1}, false}
	res7 := pawn7.isMoveValid(2, 2)
	if res7 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res6), strconv.FormatBool(false))
	}
	chessBoard.board[2][2].isEmpty = false
	chessBoard.board[2][2].piece = &pawn{pieceEntity{BLACK, 2, 2}, false}
	pawn8 := pawn{pieceEntity{WHITE, 1, 1}, false}
	res8 := pawn8.isMoveValid(2, 2)
	if res8 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res6), strconv.FormatBool(true))
	}
	pawn9 := pawn{pieceEntity{BLACK, 1, 1}, false}
	res9 := pawn9.isMoveValid(2, 2)
	if res9 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res6), strconv.FormatBool(false))
	}
}

func TestPawnDisplay(t *testing.T) {
	chessBoard = newBoard()
	pawn1 := pawn{pieceEntity{WHITE, 0, 0}, false}
	res1 := pawn1.display()
	if res1 != "WP" {
		t.Errorf("Display was incorrect, got: %s, want: %s", res1, "WP")
	}
	pawn2 := pawn{pieceEntity{BLACK, 0, 0}, false}
	res2 := pawn2.display()
	if res2 != "BP" {
		t.Errorf("Display was incorrect, got: %s, want: %s.", res2, "BP")
	}
}
