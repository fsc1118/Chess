package main

import "fmt"

const BOARD_SIZE = 8

type board struct {
	board [][]grid
}

func (b *board) printBoard() {
	// print the pieces on the board, with the row and column number
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if b.board[i][j].isEmpty {
				fmt.Print(" ")
				continue
			}
			fmt.Printf("%s ", b.board[i][j].piece.display())
		}
		fmt.Println()
	}
}

func newBoard() *board {
	initializeBoard := func() *board {
		newBoard := make([][]grid, BOARD_SIZE)
		for i := 0; i < BOARD_SIZE; i++ {
			newBoard[i] = make([]grid, BOARD_SIZE)
		}
		return &board{board: newBoard}
	}

	board := initializeBoard()
	// set all grids to empty
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			board.board[i][j].isEmpty = true
		}
	}
	return board
}

func placePiece(board *board) {
	// place white pieces
	for i := 0; i < BOARD_SIZE; i++ {
		board.board[1][i].piece = &pawn{pieceEntity: pieceEntity{color: WHITE, row: 1, column: i}}
		board.board[1][i].isEmpty = false
	}
	board.board[0][0].piece = &rock{pieceEntity: pieceEntity{color: WHITE, row: 0, column: 0}}
	board.board[0][0].isEmpty = false
	board.board[0][1].piece = &knight{pieceEntity: pieceEntity{color: WHITE, row: 0, column: 1}}
	board.board[0][1].isEmpty = false
	board.board[0][2].piece = &bishop{pieceEntity: pieceEntity{color: WHITE, row: 0, column: 2}}
	board.board[0][2].isEmpty = false
	board.board[0][3].piece = &queen{pieceEntity: pieceEntity{color: WHITE, row: 0, column: 3}}
	board.board[0][3].isEmpty = false
	board.board[0][4].piece = &king{pieceEntity: pieceEntity{color: WHITE, row: 0, column: 4}}
	board.board[0][4].isEmpty = false
	board.board[0][5].piece = &bishop{pieceEntity: pieceEntity{color: WHITE, row: 0, column: 5}}
	board.board[0][5].isEmpty = false
	board.board[0][6].piece = &knight{pieceEntity: pieceEntity{color: WHITE, row: 0, column: 6}}
	board.board[0][6].isEmpty = false
	board.board[0][7].piece = &rock{pieceEntity: pieceEntity{color: WHITE, row: 0, column: 7}}
	board.board[0][7].isEmpty = false
	// place black pieces
	for i := 0; i < BOARD_SIZE; i++ {
		board.board[6][i].piece = &pawn{pieceEntity: pieceEntity{color: BLACK, row: 6, column: i}}
		board.board[6][i].isEmpty = false
	}
	board.board[7][0].piece = &rock{pieceEntity: pieceEntity{color: BLACK, row: 7, column: 0}}
	board.board[7][0].isEmpty = false
	board.board[7][1].piece = &knight{pieceEntity: pieceEntity{color: BLACK, row: 7, column: 1}}
	board.board[7][1].isEmpty = false
	board.board[7][2].piece = &bishop{pieceEntity: pieceEntity{color: BLACK, row: 7, column: 2}}
	board.board[7][2].isEmpty = false
	board.board[7][3].piece = &queen{pieceEntity: pieceEntity{color: BLACK, row: 7, column: 3}}
	board.board[7][3].isEmpty = false
	board.board[7][4].piece = &king{pieceEntity: pieceEntity{color: BLACK, row: 7, column: 4}}
	board.board[7][4].isEmpty = false
	board.board[7][5].piece = &bishop{pieceEntity: pieceEntity{color: BLACK, row: 7, column: 5}}
	board.board[7][5].isEmpty = false
	board.board[7][6].piece = &knight{pieceEntity: pieceEntity{color: BLACK, row: 7, column: 6}}
	board.board[7][6].isEmpty = false
	board.board[7][7].piece = &rock{pieceEntity: pieceEntity{color: BLACK, row: 7, column: 7}}
	board.board[7][7].isEmpty = false
}
