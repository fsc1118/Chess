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
	// TODO
	return true
}
