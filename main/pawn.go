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
	verticalDist := abs(current.getRow() - newRow)
	horizontalDist := abs(current.getColumn() - newColumn)

	// cannot stay in the same place
	if verticalDist == 0 && horizontalDist == 0 {
		return false
	}
	// only first move can move 2 spaces
	if verticalDist == 2 && (!current.isFirstMove || horizontalDist != 0) {
		return false
	}
	// can't move more than 2 spaces forward
	if verticalDist > 2 {
		return false
	}
	// can't move more than 1 space sideways

	if horizontalDist > 1 {
		return false
	}
	if current.color == WHITE {
		// white piece can only move up
		if newRow < current.getRow() {
			return false
		}
	} else {
		// black piece can only move down
		if newRow > current.getRow() {
			return false
		}
	}

	// check if there is a piece in the way of pawn's movement
	if horizontalDist == 0 {
		// moving forward
		// check if the first space is occupied
		if current.color == WHITE {
			if !chessBoard.board[current.getRow()+1][current.getColumn()].isEmpty {
				return false
			}
			if verticalDist == 2 {
				// check if the second space is occupied
				if !chessBoard.board[current.getRow()+2][current.getColumn()].isEmpty {
					return false
				}
			}
		} else {
			if !chessBoard.board[current.getRow()-1][current.getColumn()].isEmpty {
				return false
			}
			if verticalDist == 2 {
				// check if the second space is occupied
				if !chessBoard.board[current.getRow()-2][current.getColumn()].isEmpty {
					return false
				}
			}
		}
	} else {
		// target location is not empty and must be occupied by opponent
		if chessBoard.board[newRow][newColumn].isEmpty {
			return false
		}
		if chessBoard.board[newRow][newColumn].piece.getColor() == current.color {
			return false
		}
	}
	return true
}
