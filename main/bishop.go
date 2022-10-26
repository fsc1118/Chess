/*
	A bishop is a piece that can move diagonally.
*/

package main

type bishop struct {
	pieceEntity
}

func (current *bishop) display() string {
	if current.color == WHITE {
		return "WB"
	}
	return "BB"
}

func (current *bishop) isMoveValid(newRow int, newColumn int) bool {
	if newRow < 0 || newColumn < 0 || newRow >= BOARD_SIZE || newColumn >= BOARD_SIZE {
		return false
	}
	if current.isLocationOccupiedBySelf(newRow, newColumn) {
		return false
	}
	verticalDist := abs(current.getRow() - newRow)
	horizontalDist := abs(current.getColumn() - newColumn)
	// verticalDist and horizontalDist must be equal and not 0
	if (verticalDist == 0 && horizontalDist == 0) || (horizontalDist != verticalDist) {
		return false
	}
	// check if there is a piece in the way of bishop's movement
	if newRow > current.getRow() {
		// moving down
		if newColumn > current.getColumn() {
			// moving right
			for i := 1; i < verticalDist; i++ {
				if !chessBoard.board[current.getRow()+i][current.getColumn()+i].isEmpty {
					return false
				}
			}
		} else {
			// moving left
			for i := 1; i < verticalDist; i++ {
				if !chessBoard.board[current.getRow()+i][current.getColumn()-i].isEmpty {
					return false
				}
			}
		}
	} else {
		// moving up
		if newColumn > current.getColumn() {
			// moving right
			for i := 1; i < verticalDist; i++ {
				if !chessBoard.board[current.getRow()-i][current.getColumn()+i].isEmpty {
					return false
				}
			}
		} else {
			// moving left
			for i := 1; i < verticalDist; i++ {
				if !chessBoard.board[current.getRow()-i][current.getColumn()-i].isEmpty {
					return false
				}
			}
		}
	}
	return true
}
