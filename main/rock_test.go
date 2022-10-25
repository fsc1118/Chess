package main

import (
	"testing"
)

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
