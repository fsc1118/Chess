package main

type knight struct {
	pieceEntity
}

func (current *knight) isMoveValid(newRow int, newColumn int) bool {
	// check if out of boundary
	if newRow < 0 || newColumn < 0 || newRow >= BOARD_SIZE || newColumn >= BOARD_SIZE {
		return false
	}
	isMoveValidForKnight := func() bool {
		currentRow := current.getRow()
		currentColumn := current.getColumn()
		verticalTravelDist := abs(currentRow - newRow)
		horizontalTravelDist := abs(currentColumn - newColumn)
		return (verticalTravelDist == 2 && horizontalTravelDist == 1) || (verticalTravelDist == 1 && horizontalTravelDist == 2)
	}
	return isMoveValidForKnight() && !current.isLocationOccupiedBySelf(newRow, newColumn)
}

func (current *knight) display() string {
	if current.getColor() == WHITE {
		return "WN"
	}
	return "BN"
}
