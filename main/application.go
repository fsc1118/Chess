package main

import (
	"bufio"
	"fmt"
	"os"
)

var chessBoard *board

func main() {
	chessBoard = newBoard()
	// place pieces on the board
	placePiece(chessBoard)

	// print the board
	chessBoard.printBoard()

	// Procedure:
	// do the following for each turn until the game is over (checkmate or stalemate)
	// read input from user in the form of "a2 a4"
	// check if the input is valid
	// move the piece
	// print the board

	// check if input is in the form of "a2 a4"
	isValidInput := func(input string) bool {
		if len(input) != 5 {
			return false
		}
		// check if the first character is a letter between a and h
		if input[0] < 'a' || input[0] > 'h' {
			return false
		}
		// check if the second character is a number between 1 and 8
		if input[1] < '1' || input[1] > '8' {
			return false
		}
		// check if the third character is a space
		if input[2] != ' ' {
			return false
		}
		// check if the fourth character is a letter between a and h
		if input[3] < 'a' || input[3] > 'h' {
			return false
		}
		// check if the fifth character is a number between 1 and 8
		if input[4] < '1' || input[4] > '8' {
			return false
		}
		return true
	}

	for {
		in := bufio.NewReader(os.Stdin)
		// read input from user until the user hits enter
		readed, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		// strip unprintable characters
		readed = readed[0 : len(readed)-2]

		// check if the input is valid
		if !isValidInput(readed) {
			fmt.Println("Invalid input: Input must be in the form of \"a2 a4\"")
			continue
		}
		// check if the position is empty
		if chessBoard.board[readed[1]-'1'][readed[0]-'a'].isEmpty {
			fmt.Println("Invalid input: The position at " + string(readed[0]) + string(readed[1]) + " is empty")
			continue
		}
		// check if the piece can move to the new position
		if !chessBoard.board[readed[1]-'1'][readed[0]-'a'].piece.isMoveValid(int(readed[3]-'1'), int(readed[4]-'a')) {
			fmt.Println("Invalid input: The piece at " + string(readed[0]) + string(readed[1]) + " cannot move to " + string(readed[3]) + string(readed[4]))
			continue
		}
		// move the piece
		chessBoard.board[readed[3]-'1'][readed[4]-'a'].piece.move(int(readed[1]-'1'), int(readed[0]-'a'))
	}
}
