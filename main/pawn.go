package main

type pawn struct {
	pieceEntity
	isFirstMove bool
}

func (current *pawn) display() string {
	if current.color == WHITE {
		return "WP"
	}
	return "BP"
}

func (current *pawn) isMoveValid(newRow int, newColumn int) bool {
	if newRow < 0 || newColumn < 0 || newRow >= BOARD_SIZE || newColumn >= BOARD_SIZE {
		return false
	}

	if current.isLocationOccupiedBySelf(newRow, newColumn) {
		return false
	}

	// can't go back
	if newRow <= current.row {
		return false
	}

	// can't go over 3 steps forward
	if newRow-current.row >= 3 {
		return false
	}

	// can't move two step either to left or right
	if abs(newColumn-current.column) > 1 {
		return false
	}

	// can't go two steps forward except first move
	if newRow-2 == current.row && (!current.isFirstMove || abs(newColumn-current.column) >= 1) {
		return false
	}
	// can only move horizontally if killing enemy piece
	if (newColumn == current.column+1 || newColumn == current.column-1) && chessBoard.board[newRow][newColumn].isEmpty {
		return false
	}
	return true
}
