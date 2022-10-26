package main

import (
	"strconv"
	"testing"
)

func TestPawnMoveValid(t *testing.T) {
	chessBoard = newBoard()
	// white pawn can only move forward
	pawn1 := pawn{pieceEntity{WHITE, 0, 0}, false}
	res1 := pawn1.isMoveValid(1, 0)
	if res1 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res1), strconv.FormatBool(true))
	}
	// white pawn can't move backwards
	pawn2 := pawn{pieceEntity{WHITE, 1, 0}, false}
	res2 := pawn2.isMoveValid(0, 0)
	if res2 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res2), strconv.FormatBool(false))
	}

	// back pawn can only move backwards
	pawn3 := pawn{pieceEntity{BLACK, 1, 0}, false}
	res3 := pawn3.isMoveValid(0, 0)
	if res3 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res3), strconv.FormatBool(true))
	}

	// black pawn can't move forward
	pawn4 := pawn{pieceEntity{BLACK, 0, 0}, false}
	res4 := pawn4.isMoveValid(1, 0)
	if res4 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res4), strconv.FormatBool(false))
	}

	// pawn can't move vetically for move for more than 2
	pawn5 := pawn{pieceEntity{WHITE, 0, 0}, false}
	res5 := pawn5.isMoveValid(3, 0)
	if res5 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res5), strconv.FormatBool(false))
	}

	// pawn can move forward 2 spaces if it's the first move
	pawn6 := pawn{pieceEntity{WHITE, 0, 0}, true}
	res6 := pawn6.isMoveValid(2, 0)
	if res6 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res6), strconv.FormatBool(true))
	}

	// pawn can't move forward 2 spaces if it's not the first move
	pawn7 := pawn{pieceEntity{WHITE, 0, 0}, false}
	res7 := pawn7.isMoveValid(2, 0)
	if res7 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res7), strconv.FormatBool(false))
	}

	// pawn can move forward 1 space
	pawn8 := pawn{pieceEntity{WHITE, 0, 0}, false}
	res8 := pawn8.isMoveValid(1, 0)
	if res8 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res8), strconv.FormatBool(true))
	}

	// pawn can move diagonally 1 space forward if there is an enemy piece at destination
	pawn9 := pawn{pieceEntity{WHITE, 0, 0}, false}
	chessBoard.board[1][1].isEmpty = false
	enemypiece := pieceEntity{BLACK, 1, 1}
	chessBoard.board[1][1].piece = &enemypiece
	res9 := pawn9.isMoveValid(1, 1)
	if res9 != true {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res9), strconv.FormatBool(true))
	}

	// pawn can't move diagonally 1 space forward if there is no enemy piece at destination
	pawn10 := pawn{pieceEntity{WHITE, 0, 0}, false}
	chessBoard.board[1][1].isEmpty = false
	enemypiece = pieceEntity{WHITE, 1, 1}
	chessBoard.board[1][1].piece = &enemypiece
	res10 := pawn10.isMoveValid(1, 1)
	if res10 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res10), strconv.FormatBool(false))
	}

	// there is no piece on the way to the destination
	pawn11 := pawn{pieceEntity{WHITE, 0, 0}, false}
	chessBoard.board[1][0].isEmpty = false
	chessBoard.board[1][0].piece = &enemypiece
	res11 := pawn11.isMoveValid(1, 0)
	if res11 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res11), strconv.FormatBool(false))
	}
	// new board
	chessBoard = newBoard()
	pawn12 := pawn{pieceEntity{WHITE, 0, 0}, true}
	chessBoard.board[2][0].isEmpty = false
	chessBoard.board[2][0].piece = &queen{pieceEntity{BLACK, 1, 0}}
	res12 := pawn12.isMoveValid(2, 0)
	if res12 != false {
		t.Errorf("Move check was incorrect, got: %s, want: %s", strconv.FormatBool(res12), strconv.FormatBool(false))
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
