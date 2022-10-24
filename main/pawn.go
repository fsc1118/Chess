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
	// TODO
	return false
}
