package main

import (
	"strconv"
	"testing"
)

func TestKnightMoveValid(t *testing.T) {
	chessBoard = newBoard()
	knight1 := knight{pieceEntity{BLACK, 0, 0}}
	res1 := knight1.isMoveValid(0, 1)
	if res1 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s.", strconv.FormatBool(res1), strconv.FormatBool(false))
	}
	knight2 := knight{pieceEntity{WHITE, 0, 0}}
	res2 := knight2.isMoveValid(2, 1)
	if res2 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res2), strconv.FormatBool(true))
	}
	knight3 := knight{pieceEntity{WHITE, 0, 0}}
	res3 := knight3.isMoveValid(0, -1)
	if res3 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res3), strconv.FormatBool(false))
	}
	knight4 := knight{pieceEntity{WHITE, 3, 3}}
	res4 := knight4.isMoveValid(1, 2)
	if res4 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res4), strconv.FormatBool(true))
	}
}

func TestKnightDisplay(t *testing.T) {
	chessBoard = newBoard()
	knight1 := knight{pieceEntity{WHITE, 0, 0}}
	res1 := knight1.display()
	if res1 != "WN" {
		t.Errorf("Display was incorrect, got: %s, want: %s", res1, "WN")
	}
	knight2 := knight{pieceEntity{BLACK, 0, 0}}
	res2 := knight2.display()
	if res2 != "BN" {
		t.Errorf("Display was incorrect, got: %s, want: %s.", res2, "BN")
	}
}
