package main

import (
	"strconv"
	"testing"
)

func TestBishopDisplay(t *testing.T) {
	chessBoard = newBoard()
	bishop1 := bishop{pieceEntity{WHITE, 0, 0}}
	res1 := bishop1.display()
	if res1 != "WB" {
		t.Errorf("Display was incorrect, got: %s, want: %s", res1, "WB")
	}
	bishop2 := bishop{pieceEntity{BLACK, 0, 0}}
	res2 := bishop2.display()
	if res2 != "BB" {
		t.Errorf("Display was incorrect, got: %s, want: %s.", res2, "BB")
	}
}

func TestBishopValidMove(t *testing.T) {
	chessBoard = newBoard()
	bishop1 := bishop{pieceEntity{BLACK, 0, 0}}
	res1 := bishop1.isMoveValid(0, 1)
	if res1 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s.", strconv.FormatBool(res1), strconv.FormatBool(false))
	}
	bishop2 := bishop{pieceEntity{WHITE, 0, 0}}
	res2 := bishop2.isMoveValid(2, 2)
	if res2 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res2), strconv.FormatBool(true))
	}
	bishop3 := bishop{pieceEntity{WHITE, 0, 0}}
	res3 := bishop3.isMoveValid(0, -1)
	if res3 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res3), strconv.FormatBool(false))
	}
	bishop4 := bishop{pieceEntity{WHITE, 3, 3}}
	res4 := bishop4.isMoveValid(1, 1)
	if res4 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res4), strconv.FormatBool(true))
	}

	// bishop is blocked by another piece
	bishop5 := bishop{pieceEntity{WHITE, 3, 3}}
	chessBoard.board[2][2].isEmpty = false
	res5 := bishop5.isMoveValid(1, 1)

	if res5 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res5), strconv.FormatBool(false))
	}
}
