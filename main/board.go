package main

const BOARD_SIZE = 8

type board struct {
	board [][]grid
}

func newBoard() *board {
	initializeBoard := func() *board {
		newBoard := make([][]grid, BOARD_SIZE)
		for i := 0; i < BOARD_SIZE; i++ {
			newBoard[i] = make([]grid, BOARD_SIZE)
		}
		return &board{board: newBoard}
	}
	placePiece := func(board *board) {
		for i := 0; i < BOARD_SIZE; i++ {
			for j := 0; j < BOARD_SIZE; j++ {
				knight := knight{pieceEntity{color: WHITE, row: i, column: i}}
				newGrid := grid{isEmpty: true, piece: &knight}
				board.board[i][j] = newGrid
			}
		}
	}
	board := initializeBoard()
	placePiece(board)
	return board
}
