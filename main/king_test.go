package main

import (
	"strconv"
	"testing"
)

func TestKingMoveValid(t *testing.T) {
	chessBoard = newBoard()
	king1 := king{pieceEntity{BLACK, 0, 0}}
	res1 := king1.isMoveValid(0, 2)
	if res1 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s.", strconv.FormatBool(res1), strconv.FormatBool(false))
	}
	king2 := king{pieceEntity{WHITE, 0, 0}}
	res2 := king2.isMoveValid(0, 1)
	if res2 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res2), strconv.FormatBool(true))
	}
	king3 := king{pieceEntity{WHITE, 0, 0}}
	res3 := king3.isMoveValid(0, -1)
	if res3 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res3), strconv.FormatBool(false))
	}
	king4 := king{pieceEntity{WHITE, 0, 0}}
	res4 := king4.isMoveValid(1, 1)
	if res4 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res4), strconv.FormatBool(true))
	}
}

func TestKingDisplay(t *testing.T) {
	chessBoard = newBoard()
	king1 := king{pieceEntity{WHITE, 0, 0}}
	res1 := king1.display()
	if res1 != "WK" {
		t.Errorf("Display was incorrect, got: %s, want: %s", res1, "WK")
	}
	king2 := king{pieceEntity{BLACK, 0, 0}}
	res2 := king2.display()
	if res2 != "BK" {
		t.Errorf("Display was incorrect, got: %s, want: %s.", res2, "BK")
	}
}
