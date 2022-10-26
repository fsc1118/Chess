package main

import (
	"strconv"
	"testing"
)

func TestQueenDisplay(t *testing.T) {
	chessBoard = newBoard()
	queen1 := queen{pieceEntity{WHITE, 0, 0}}
	res1 := queen1.display()
	if res1 != "WQ" {
		t.Errorf("Display was incorrect, got: %s, want: %s", res1, "WQ")
	}
	queen2 := queen{pieceEntity{BLACK, 0, 0}}
	res2 := queen2.display()
	if res2 != "BQ" {
		t.Errorf("Display was incorrect, got: %s, want: %s.", res2, "BQ")
	}
}

func TestQueenValidMove(t *testing.T) {
	chessBoard = newBoard()
	// check queen move horizontally
	queen1 := queen{pieceEntity{BLACK, 0, 0}}
	res1 := queen1.isMoveValid(0, 1)
	if res1 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s.", strconv.FormatBool(res1), strconv.FormatBool(true))
	}
	// check queen move diagonally
	queen2 := queen{pieceEntity{WHITE, 0, 0}}
	res2 := queen2.isMoveValid(2, 2)
	if res2 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res2), strconv.FormatBool(true))
	}
	// check queen moves out of bounds
	queen3 := queen{pieceEntity{WHITE, 0, 0}}
	res3 := queen3.isMoveValid(0, -1)
	if res3 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res3), strconv.FormatBool(false))
	}
	// check queen moves vertically
	queen4 := queen{pieceEntity{WHITE, 3, 3}}
	res4 := queen4.isMoveValid(4, 3)
	if res4 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res4), strconv.FormatBool(true))
	}

	// check quuen is blocked by another piece diagonally
	queen5 := queen{pieceEntity{WHITE, 3, 3}}
	chessBoard.board[2][2].isEmpty = false
	res5 := queen5.isMoveValid(1, 1)
	if res5 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res5), strconv.FormatBool(false))
	}

	// check queen is blocked by another piece horizontally
	queen6 := queen{pieceEntity{WHITE, 3, 3}}
	chessBoard.board[3][2].isEmpty = false
	res6 := queen6.isMoveValid(3, 1)
	if res6 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res6), strconv.FormatBool(false))
	}

	// check queen is blocked by another piece vertically
	queen7 := queen{pieceEntity{WHITE, 3, 3}}
	chessBoard.board[2][3].isEmpty = false
	res7 := queen7.isMoveValid(1, 3)
	if res7 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res7), strconv.FormatBool(false))
	}

	// renew the board
	chessBoard = newBoard()
	// queen can't move with both x and y coordinates changed while the change in x and y are not equal
	queen8 := queen{pieceEntity{WHITE, 3, 3}}
	res8 := queen8.isMoveValid(4, 5)
	if res8 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res8), strconv.FormatBool(false))
	}
}
