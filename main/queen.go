// queen is a piece that can move in horizontally, vertically, and diagonally for infinite distance

package main

type queen struct {
	pieceEntity
}

func (current *queen) display() string {
	if current.color == WHITE {
		return "WQ"
	}
	return "BQ"
}

func (current *queen) isMoveValid(newRow int, newColumn int) bool {
	if newRow < 0 || newColumn < 0 || newRow >= BOARD_SIZE || newColumn >= BOARD_SIZE {
		return false
	}
	if current.isLocationOccupiedBySelf(newRow, newColumn) {
		return false
	}
	verticalDist := abs(current.getRow() - newRow)
	horizontalDist := abs(current.getColumn() - newColumn)
	// either verticalDist isn't 0 or horizontalDist isn't 0, or verticalDist == horizontalDist (diagonal), but can't all three

	if (verticalDist == 0 && horizontalDist == 0) || (horizontalDist != 0 && verticalDist != 0 && verticalDist != horizontalDist) {
		return false
	}
	// check if there is a piece in the way (if moving diagonally)
	if verticalDist == horizontalDist {
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
	} else {
		// check if there is a piece in the way (if moving vertically or horizontally)
		if verticalDist != 0 {
			// moving vertically
			if newRow > current.getRow() {
				// moving down
				for i := current.getRow() + 1; i < newRow; i++ {
					if !chessBoard.board[i][current.getColumn()].isEmpty {
						return false
					}
				}
			} else {
				// moving up
				for i := current.getRow() - 1; i > newRow; i-- {
					if !chessBoard.board[i][current.getColumn()].isEmpty {
						return false
					}
				}
			}
		} else {
			// moving horizontally
			if newColumn > current.getColumn() {
				// moving right
				for i := current.getColumn() + 1; i < newColumn; i++ {
					if !chessBoard.board[current.getRow()][i].isEmpty {
						return false
					}
				}
			} else {
				// moving left
				for i := current.getColumn() - 1; i > newColumn; i-- {
					if !chessBoard.board[current.getRow()][i].isEmpty {
						return false
					}
				}
			}
		}
	}

	return true
}
