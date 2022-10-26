package main

type rock struct {
	pieceEntity
}

func (current *rock) display() string {
	if current.color == WHITE {
		return "WR"
	}
	return "BR"
}

func (current *rock) isMoveValid(newRow int, newColumn int) bool {
	if newRow < 0 || newColumn < 0 || newRow >= BOARD_SIZE || newColumn >= BOARD_SIZE {
		return false
	}
	if current.isLocationOccupiedBySelf(newRow, newColumn) {
		return false
	}
	verticalDist := abs(current.getRow() - newRow)
	horizontalDist := abs(current.getColumn() - newColumn)
	// either verticalDist isn't 0 or horizontalDist isn't 0, but can't both
	if (verticalDist == 0 && horizontalDist == 0) || (horizontalDist != 0 && verticalDist != 0) {
		return false
	}
	// check if there is a piece in the way
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
	return true
}
