package main

type king struct {
	pieceEntity
}

func (current king) display() string {
	if current.getColor() == WHITE {
		return "WK"
	}
	return "BK"
}

func (current king) isMoveValid(newRow int, newColumn int) bool {
	// check if out of boundary
	if newRow < 0 || newColumn < 0 || newRow >= BOARD_SIZE || newColumn >= BOARD_SIZE {
		return false
	}
	isMoveValidForKing := func() bool {
		currentRow := current.getRow()
		currentColumn := current.getColumn()
		verticalTravelDist := abs(currentRow - newRow)
		horizontalTravelDist := abs(currentColumn - newColumn)
		return (verticalTravelDist <= 1 && horizontalTravelDist <= 1) && (verticalTravelDist != 0 || horizontalTravelDist != 0)
	}
	return isMoveValidForKing() && !current.isLocationOccupiedBySelf(newRow, newColumn)
}
