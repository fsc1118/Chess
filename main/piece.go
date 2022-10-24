package main

const BLACK = 1
const WHITE = 2

type color int

type piece interface {
	locatable
	displayable
	moveable
	getColor() color
}

type locatable interface {
	getRow() int
	getColumn() int
	setRow(newRow int)
	setColumn(newColumn int)
	isLocationOccupiedBySelf(newRow int, newColumn int) bool
}

type displayable interface {
	display() string
}

type moveable interface {
	move(newRow int, newColumn int)
	isMoveValid(newRow int, newColumn int) bool
}

type pieceEntity struct {
	color  color
	row    int
	column int
}

func (current *pieceEntity) getRow() int                                { return current.row }
func (current *pieceEntity) getColumn() int                             { return current.column }
func (current *pieceEntity) getColor() color                            { return current.color }
func (current *pieceEntity) setRow(newRow int)                          { current.row = newRow }
func (current *pieceEntity) setColumn(newColumn int)                    { current.column = newColumn }
func (current *pieceEntity) display() string                            { return "" }
func (current *pieceEntity) isMoveValid(newRow int, newColumn int) bool { return false }
func (current *pieceEntity) move(newRow int, newColumn int) {
	current.setColumn(newColumn)
	current.setRow(newRow)
	chessBoard.board[newRow][newColumn].isEmpty = false
	chessBoard.board[newRow][newColumn].piece = current
}
func (current *pieceEntity) isLocationOccupiedBySelf(newRow int, newColumn int) bool {
	destination := chessBoard.board[newRow][newColumn]
	if destination.isEmpty {
		return false
	}
	return destination.piece.getColor() != current.getColor()
}
